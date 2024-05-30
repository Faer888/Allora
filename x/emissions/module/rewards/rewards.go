package rewards

import (
	"fmt"

	"cosmossdk.io/errors"
	cosmosMath "cosmossdk.io/math"
	"github.com/allora-network/allora-chain/app/params"
	alloraMath "github.com/allora-network/allora-chain/math"
	"github.com/allora-network/allora-chain/x/emissions/keeper"
	"github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EmitRewards(
	ctx sdk.Context,
	k keeper.Keeper,
	blockHeight BlockHeight,
	weights map[uint64]*alloraMath.Dec,
	sumWeight alloraMath.Dec,
	totalRevenue cosmosMath.Int,
) error {
	totalReward, err := k.GetTotalRewardToDistribute(ctx)
	ctx.Logger().Debug(fmt.Sprintf("Reward to distribute this epoch: %s", totalReward.String()))
	if err != nil {
		return errors.Wrapf(err, "failed to get total reward to distribute")
	}
	if totalReward.IsZero() {
		ctx.Logger().Warn("The total scheduled rewards to distribute this epoch are zero!")
		return nil
	}

	moduleParams, err := k.GetParams(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed to get module params")
	}

	sortedTopics := alloraMath.GetSortedKeys(weights)

	// Distribute rewards between topics
	topicRewards, err := GenerateRewardsDistributionByTopic(
		ctx,
		k,
		moduleParams.MaxTopicsPerBlock,
		blockHeight,
		totalReward,
		weights,
		sortedTopics,
		sumWeight,
		totalRevenue,
	)
	if err != nil {
		return errors.Wrapf(err, "failed to generate total reward by topic")
		// Will return nil if there are no topics to reward
	} else if topicRewards == nil {
		return nil
	}

	totalRewardToStakedReputers := alloraMath.ZeroDec()
	// for every topic
	for _, topicId := range sortedTopics {
		topicReward := topicRewards[topicId]
		if topicReward == nil {
			ctx.Logger().Warn(fmt.Sprintf("Topic %d has no reward, skipping", topicId))
			continue
		}
		// Get topic reward nonce/block height
		topicRewardNonce, err := k.GetTopicRewardNonce(ctx, topicId)
		// If the topic has no reward nonce, skip it
		if err != nil || topicRewardNonce == 0 {
			continue
		}

		// Distribute rewards between topic participants
		totalRewardsDistribution, rewardInTopicToReputers, err := GenerateRewardsDistributionByTopicParticipant(
			ctx,
			k,
			topicId,
			topicReward,
			topicRewardNonce,
			moduleParams,
		)
		if err != nil {
			topicRewardString := "nil"
			ctx.Logger().Warn(
				fmt.Sprintf(
					"Failed to Generate Rewards for Topic, Skipping:\nTopic Id %d\nTopic Reward Amount %s\nError:\n%s\n\n",
					topicId,
					topicRewardString,
					err.Error(),
				),
			)
			continue
		}
		totalRewardToStakedReputers, err = totalRewardToStakedReputers.Add(rewardInTopicToReputers)
		if err != nil {
			return errors.Wrapf(
				err,
				"Error finding sum of rewards to Reputers:\n%s\n%s",
				totalRewardToStakedReputers.String(),
				rewardInTopicToReputers.String(),
			)
		}

		// Pay out rewards to topic participants
		payoutErrors := payoutRewards(ctx, k, totalRewardsDistribution)
		if len(payoutErrors) > 0 {
			for _, err := range payoutErrors {
				ctx.Logger().Warn(
					fmt.Sprintf(
						"Failed to pay out rewards to participant in Topic:\nTopic Id %d\nTopic Reward Amount %s\nError:\n%s\n\n",
						topicId,
						topicReward.String(),
						err.Error(),
					),
				)
			}
			continue
		}

		// Prune records after rewards have been paid out
		err = pruneRecordsAfterRewards(
			ctx,
			k,
			moduleParams.MinEpochLengthRecordLimit,
			topicId,
			topicRewardNonce,
		)
		if err != nil {
			ctx.Logger().Warn(
				fmt.Sprintf(
					"Failed to prune records after rewards for Topic, Skipping:\nTopic Id %d\nTopic Reward Amount %s\nError:\n%s\n\n",
					topicId,
					topicReward.String(),
					err.Error(),
				),
			)
			continue
		}
	}
	ctx.Logger().Debug(
		fmt.Sprintf("Paid out %s to staked reputers over %d topics",
			totalRewardToStakedReputers.String(),
			len(topicRewards)))
	if !totalReward.IsZero() && uint64(blockHeight)%moduleParams.BlocksPerMonth == 0 {
		// set the previous percentage reward to staked reputers
		// for the mint module to be able to control the inflation rate to that actor
		percentageToStakedReputers, err := totalRewardToStakedReputers.Quo(totalReward)
		if err != nil {
			return errors.Wrapf(err, "failed to calculate percentage to staked reputers")
		}
		err = k.SetPreviousPercentageRewardToStakedReputers(ctx, percentageToStakedReputers)
		if err != nil {
			return errors.Wrapf(err, "failed to set previous percentage reward to staked reputers")
		}
	}

	return nil
}

func GenerateRewardsDistributionByTopic(
	ctx sdk.Context,
	k keeper.Keeper,
	maxTopicsPerBlock uint64,
	blockHeight BlockHeight,
	totalReward alloraMath.Dec,
	weights map[uint64]*alloraMath.Dec,
	sortedTopics []uint64,
	sumWeight alloraMath.Dec,
	totalRevenue cosmosMath.Int,
) (map[uint64]*alloraMath.Dec, error) {
	if sumWeight.IsZero() {
		ctx.Logger().Warn("No weights, no rewards!")
		return nil, nil
	}
	// Filter out topics that are not reward-ready, inactivate if needed
	// Update sum weight and revenue
	weightsOfActiveTopics, sumWeight, err := FilterAndInactivateTopicsUpdatingSums(
		ctx,
		k,
		weights,
		sortedTopics,
		sumWeight,
		totalReward,
		blockHeight,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to inactivate topics and update sums")
	}
	if sumWeight.IsZero() {
		ctx.Logger().Warn("No filtered weights, no rewards!")
		return nil, nil
	}

	// Sort remaining active topics by weight desc and skim the top via SortTopicsByReturnDescWithRandomTiebreaker() and param MaxTopicsPerBlock
	weightsOfTopActiveTopics, _ := SkimTopTopicsByWeightDesc(
		ctx,
		weightsOfActiveTopics,
		maxTopicsPerBlock,
		blockHeight,
	)

	// Return the revenue to those topics that didn't make the cut
	// Loop though sortedTopics and if the topic is not in sortedTopics, add to running revenue sum
	sumRevenueOfBottomTopics := cosmosMath.ZeroInt()
	sumWeightOfBottomTopics := alloraMath.ZeroDec()
	for _, topicId := range sortedTopics {
		// If the topic is in weightsOfActiveTopics but not in weightsOfTopActiveTopics, add its revenue to the running sum
		if _, isActive := weightsOfActiveTopics[topicId]; isActive {
			if _, isTop := weightsOfTopActiveTopics[topicId]; !isTop {
				topicFeeRevenue, err := k.GetTopicFeeRevenue(ctx, topicId)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to get topic fee revenue")
				}
				sumRevenueOfBottomTopics = sumRevenueOfBottomTopics.Add(topicFeeRevenue.Revenue)
				sumWeightOfBottomTopics, err = sumWeightOfBottomTopics.Add(*weights[topicId])
				if err != nil {
					return nil, errors.Wrapf(err, "failed to add weight to sum")
				}
			} else {
				// For topics that are active and top, they will get their rewards paid out this block.
				// Everybody else doesn't. Therefore, for topics that have topicFeeRevenue but havent received rewards,
				// we don't reset their topic fee revenue, effectively "double counting" their topic fee revenue
				//  giving them a chance to earn rewards in future blocks as they accumulate more fees.
				//
				// This call must come after GetTopicFeeRevenue() is last called per topic in
				// GetAndOptionallyUpdateActiveTopicWeights -> GetCurrentTopicWeight
				// because otherwise the returned revenue would be zero
				err = k.ResetTopicFeeRevenue(ctx, topicId, blockHeight)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to reset topic fee revenue")
				}
			}
		}
		ctx.Logger().Debug("Topic ID: ", topicId, " is not in weightsOfActiveTopics")
	}

	sortedTopTopics := alloraMath.GetSortedKeys(weightsOfTopActiveTopics)

	weightOfTopTopics, err := sumWeight.Sub(sumWeightOfBottomTopics)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to subtract weight of bottom topics from sum")
	}
	// Revenue (above) is what was earned by topics in this timestep. Rewards are what are actually paid to topics => participants
	// The reward and revenue calculations are coupled here to minimize excessive
	topicRewards, err := CalcTopicRewards(
		ctx,
		k,
		weightsOfTopActiveTopics,
		sortedTopTopics,
		weightOfTopTopics,
		totalReward,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to calculate topic rewards")
	}

	return topicRewards, nil
}

func removeFromSumWeightAndRevenue(
	sumWeight alloraMath.Dec,
	weight *alloraMath.Dec,
) (alloraMath.Dec, error) {
	// Update sum weight and revenue -- We won't be deducting fees from inactive topics, as we won't be churning them
	// i.e. we'll neither emit their worker/reputer requests or calculate rewards for its participants this epoch
	sumWeight, err := sumWeight.Sub(*weight)
	if err != nil {
		return alloraMath.Dec{}, errors.Wrapf(err, "failed to subtract weight from sum")
	}
	return sumWeight, nil
}

func FilterAndInactivateTopicsUpdatingSums(
	ctx sdk.Context,
	k keeper.Keeper,
	weights map[uint64]*alloraMath.Dec,
	sortedTopics []uint64,
	sumWeight alloraMath.Dec,
	totalReward alloraMath.Dec,
	blockHeight BlockHeight,
) (
	map[uint64]*alloraMath.Dec,
	alloraMath.Dec,
	error,
) {
	moduleParams, err := k.GetParams(ctx)
	if err != nil {
		return nil, alloraMath.Dec{}, errors.Wrapf(err, "failed to get min topic weight")
	}

	weightsOfActiveTopics := make(map[TopicId]*alloraMath.Dec)
	for _, topicId := range sortedTopics {
		weight := weights[topicId]
		// Filter out if not reward-ready
		// Check topic has an unfulfilled reward nonce
		rewardNonce, err := k.GetTopicRewardNonce(ctx, topicId)
		filterOutTopic := false
		filterOutErrorMessage := ""
		if err != nil {
			ctx.Logger().Warn(fmt.Sprintf("Error getting reputer request nonces: %s", err.Error()))
			filterOutTopic = true
			filterOutErrorMessage = "failed to remove from sum weight and revenue"
		}
		if rewardNonce == 0 {
			ctx.Logger().Warn("Reward nonce is 0")
			filterOutTopic = true
			filterOutErrorMessage = "failed to remove nil-reputer-nonce topic from sum weight and revenue"
		}

		// Inactivate and skip the topic if its weight is below the globally-set minimum
		if weight.Lt(moduleParams.MinTopicWeight) {
			ctx.Logger().Warn(fmt.Sprintf("Topic weight is below the minimum: %d", topicId))
			err = k.InactivateTopic(ctx, topicId)
			if err != nil {
				return nil, alloraMath.Dec{}, errors.Wrapf(err, "failed to inactivate topic")
			}

			// This way we won't double count from this earlier epoch revenue the next time this topic is activated
			// This must come after GetTopicFeeRevenue() is last called per topic because otherwise the returned revenue will be zero
			err = k.ResetTopicFeeRevenue(ctx, topicId, blockHeight)
			if err != nil {
				return nil, alloraMath.Dec{}, errors.Wrapf(err, "failed to reset topic fee revenue")
			}

			// Update sum weight and revenue -- We won't be deducting fees from inactive topics, as we won't be churning them
			// i.e. we'll neither emit their worker/reputer requests or calculate rewards for its participants this epoch
			filterOutTopic = true
			filterOutErrorMessage = "failed to remove inactivated from sum weight and revenue"
		}
		if filterOutTopic {
			sumWeight, err = removeFromSumWeightAndRevenue(sumWeight, weight)
			if err != nil {
				return nil, alloraMath.Dec{}, errors.Wrapf(err, filterOutErrorMessage)
			}
		} else {
			weightsOfActiveTopics[topicId] = weight
		}
	}
	return weightsOfActiveTopics, sumWeight, nil
}

func CalcTopicRewards(
	ctx sdk.Context,
	k keeper.Keeper,
	weights map[uint64]*alloraMath.Dec,
	sortedTopics []uint64,
	sumWeight alloraMath.Dec,
	totalReward alloraMath.Dec,
) (
	map[uint64]*alloraMath.Dec,
	error,
) {
	topicRewards := make(map[TopicId]*alloraMath.Dec)
	for _, topicId := range sortedTopics {
		weight := weights[topicId]
		topicRewardFraction, err := GetTopicRewardFraction(weight, sumWeight)
		if err != nil {
			return nil, errors.Wrapf(err, "topic reward fraction error")
		}
		topicReward, err := GetTopicReward(topicRewardFraction, totalReward)
		if err != nil {
			return nil, errors.Wrapf(err, "topic reward error")
		}
		topicRewards[topicId] = &topicReward
	}
	return topicRewards, nil
}

func GenerateRewardsDistributionByTopicParticipant(
	ctx sdk.Context,
	k keeper.Keeper,
	topicId uint64,
	topicReward *alloraMath.Dec,
	blockHeight int64,
	moduleParams types.Params,
) (
	totalRewardsDistribution []TaskRewards,
	taskReputerReward alloraMath.Dec,
	err error,
) {
	if topicReward == nil {
		return nil, alloraMath.Dec{}, types.ErrInvalidReward
	}
	bundles, err := k.GetReputerLossBundlesAtBlock(ctx, topicId, blockHeight)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get network loss bundle at block %d",
			blockHeight,
		)
	}

	lossBundles, err := k.GetNetworkLossBundleAtBlock(ctx, topicId, blockHeight)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get network loss bundle at block %d",
			blockHeight,
		)
	}

	// Calculate and Set the reputer scores
	reputerScores, err := GenerateReputerScores(ctx, k, topicId, blockHeight, *bundles)
	if err != nil {
		return nil, alloraMath.Dec{}, err
	}

	// Calculate and Set the worker scores for their inference work
	infererScores, err := GenerateInferenceScores(ctx, k, topicId, blockHeight, *lossBundles)
	if err != nil {
		return nil, alloraMath.Dec{}, err
	}

	// Calculate and Set the worker scores for their forecast work
	forecasterScores, err := GenerateForecastScores(ctx, k, topicId, blockHeight, *lossBundles)
	if err != nil {
		return nil, alloraMath.Dec{}, err
	}

	// Get reputer participants' addresses and reward fractions to be used in the reward round for topic
	reputers, reputersRewardFractions, err := GetReputersRewardFractions(
		ctx,
		k,
		topicId,
		moduleParams.PRewardReputer,
		reputerScores,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get reputer reward round data",
		)
	}

	// Get reputer task entropy
	reputerEntropy, err := GetReputerTaskEntropy(
		ctx,
		k,
		topicId,
		moduleParams.TaskRewardAlpha,
		moduleParams.BetaEntropy,
		reputers,
		reputersRewardFractions,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get reputer task entropy",
		)
	}

	// Get inferer reward fractions
	inferers, inferersRewardFractions, err := GetInferenceTaskRewardFractions(
		ctx,
		k,
		topicId,
		blockHeight,
		moduleParams.PRewardInference,
		moduleParams.CRewardInference,
		infererScores,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get inferer reward fractions",
		)
	}

	// Get inference entropy
	inferenceEntropy, err := GetInferenceTaskEntropy(
		ctx,
		k,
		topicId,
		moduleParams.TaskRewardAlpha,
		moduleParams.BetaEntropy,
		inferers,
		inferersRewardFractions,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get inference task entropy",
		)
	}

	// Get forecaster reward fractions
	forecasters, forecastersRewardFractions, err := GetForecastingTaskRewardFractions(
		ctx,
		k,
		topicId,
		blockHeight,
		moduleParams.PRewardForecast,
		moduleParams.CRewardForecast,
		forecasterScores,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get forecaster reward fractions",
		)
	}

	var forecastingEntropy alloraMath.Dec
	if len(forecasters) > 0 && len(inferers) > 1 {
		// Get forecasting entropy
		forecastingEntropy, err = GetForecastingTaskEntropy(
			ctx,
			k,
			topicId,
			moduleParams.TaskRewardAlpha,
			moduleParams.BetaEntropy,
			forecasters,
			forecastersRewardFractions,
		)
		if err != nil {
			return []TaskRewards{}, alloraMath.Dec{}, err
		}
	} else {
		// If there are no forecasters, set forecasting entropy to zero
		forecastingEntropy = alloraMath.ZeroDec()
	}

	// Get Total Rewards for Reputation task
	taskReputerReward, err = GetRewardForReputerTaskInTopic(
		inferenceEntropy,
		forecastingEntropy,
		reputerEntropy,
		topicReward,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get reward for reputer task in topic",
		)
	}

	// Get Total Rewards for Inference task
	taskInferenceReward, err := GetRewardForInferenceTaskInTopic(
		lossBundles.NaiveValue,
		lossBundles.CombinedValue,
		inferenceEntropy,
		forecastingEntropy,
		reputerEntropy,
		topicReward,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get reward for inference task in topic",
		)
	}

	// Get Total Rewards for Forecasting task
	taskForecastingReward, err := GetRewardForForecastingTaskInTopic(
		lossBundles.NaiveValue,
		lossBundles.CombinedValue,
		inferenceEntropy,
		forecastingEntropy,
		reputerEntropy,
		topicReward,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get reward for forecasting task in topic",
		)
	}

	totalRewardsDistribution = make([]TaskRewards, 0)

	// Get Distribution of Rewards per Reputer
	reputerRewards, err := GetRewardPerReputer(
		ctx,
		k,
		topicId,
		taskReputerReward,
		reputers,
		reputersRewardFractions,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(err, "failed to get reputer rewards")
	}
	totalRewardsDistribution = append(totalRewardsDistribution, reputerRewards...)

	// Get Distribution of Rewards per Worker - Inference Task
	inferenceRewards, err := GetRewardPerWorker(
		topicId,
		WorkerInferenceRewardType,
		taskInferenceReward,
		inferers,
		inferersRewardFractions,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get inference rewards",
		)
	}
	totalRewardsDistribution = append(totalRewardsDistribution, inferenceRewards...)

	// Get Distribution of Rewards per Worker - Forecast Task
	forecastRewards, err := GetRewardPerWorker(
		topicId,
		WorkerForecastRewardType,
		taskForecastingReward,
		forecasters,
		forecastersRewardFractions,
	)
	if err != nil {
		return []TaskRewards{}, alloraMath.Dec{}, errors.Wrapf(
			err,
			"failed to get forecast rewards",
		)
	}
	totalRewardsDistribution = append(totalRewardsDistribution, forecastRewards...)

	return totalRewardsDistribution, taskReputerReward, nil
}

// pay out the rewards to the participants
// this function moves tokens from the rewards module to the participants
// if it fails to pay a particular participant, it will continue to the next participant
func payoutRewards(
	ctx sdk.Context,
	k keeper.Keeper,
	rewards []TaskRewards,
) []error {
	ret := make([]error, 0)
	for _, reward := range rewards {
		if reward.Reward.IsZero() {
			continue
		}

		rewardInt := reward.Reward.Abs().SdkIntTrim()
		coins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, rewardInt))

		if reward.Type == ReputerRewardType {
			err := k.SendCoinsFromModuleToModule(
				ctx,
				types.AlloraRewardsAccountName,
				types.AlloraStakingAccountName,
				coins,
			)
			if err != nil {
				ret = append(ret, errors.Wrapf(
					err,
					"failed to send coins from rewards module to staking module: %s",
					coins.String(),
				))
				continue
			}
			err = k.AddStake(ctx, reward.TopicId, reward.Address, cosmosMath.Int(rewardInt))
			if err != nil {
				ret = append(
					ret,
					errors.Wrapf(
						err,
						"failed to add stake %s: %s",
						reward.Address,
						rewardInt.String(),
					),
				)
				continue
			}
		} else {
			accAddress, err := sdk.AccAddressFromBech32(reward.Address)
			if err != nil {
				ret = append(ret, errors.Wrapf(err, "failed to decode payout address: %s", reward.Address))
				continue
			}
			err = k.BankKeeper().SendCoinsFromModuleToAccount(
				ctx,
				types.AlloraRewardsAccountName,
				accAddress,
				coins,
			)
			if err != nil {
				ret = append(ret, errors.Wrapf(
					err,
					"failed to send coins from rewards module to payout address %s, %s",
					types.AlloraRewardsAccountName,
					reward.Address,
				))
				continue
			}
		}
	}

	return ret
}

func pruneRecordsAfterRewards(
	ctx sdk.Context,
	k keeper.Keeper,
	minEpochLengthRecordLimit int64,
	topicId uint64,
	topicRewardNonce int64,
) error {
	// Delete topic reward nonce
	err := k.DeleteTopicRewardNonce(ctx, topicId)
	if err != nil {
		return errors.Wrapf(err, "failed to delete topic reward nonce")
	}

	// Get oldest unfulfilled nonce - delete everything behind it
	unfulfilledNonces, err := k.GetUnfulfilledReputerNonces(ctx, topicId)
	if err != nil {
		return err
	}

	// Assume the oldest nonce is the topic reward nonce
	oldestNonce := topicRewardNonce
	// If there are unfulfilled nonces, find the oldest one
	if len(unfulfilledNonces.Nonces) > 0 {
		oldestNonce = unfulfilledNonces.Nonces[0].ReputerNonce.BlockHeight
		for _, nonce := range unfulfilledNonces.Nonces {
			if nonce.ReputerNonce.BlockHeight < oldestNonce {
				oldestNonce = nonce.ReputerNonce.BlockHeight
			}
		}
	}

	topic, err := k.GetTopic(ctx, topicId)
	if err != nil {
		return errors.Wrapf(err, "failed to get topic")
	}

	// Prune records x EpochsLengths behind the oldest nonce
	// This is to leave the necessary data for the remaining
	// unfulfilled nonces to be fulfilled
	oldestNonce -= minEpochLengthRecordLimit * topic.EpochLength

	// Prune old records after rewards have been paid out
	err = k.PruneRecordsAfterRewards(ctx, topicId, oldestNonce)
	if err != nil {
		return errors.Wrapf(err, "failed to prune records after rewards")
	}

	return nil
}
