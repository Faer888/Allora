package queryserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	synth "github.com/allora-network/allora-chain/x/emissions/keeper/inference_synthesis"
	"github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetWorkerLatestInferenceByTopicId handles the query for the latest inference by a specific worker for a given topic.
func (qs queryServer) GetWorkerLatestInferenceByTopicId(
	ctx context.Context,
	req *types.QueryWorkerLatestInferenceRequest,
) (*types.QueryWorkerLatestInferenceResponse, error) {
	topicExists, err := qs.k.TopicExists(ctx, req.TopicId)
	if !topicExists {
		return nil, status.Errorf(codes.NotFound, "topic %v not found", req.TopicId)
	} else if err != nil {
		return nil, err
	}

	inference, err := qs.k.GetWorkerLatestInferenceByTopicId(ctx, req.TopicId, req.WorkerAddress)
	if err != nil {
		return nil, err
	}

	return &types.QueryWorkerLatestInferenceResponse{LatestInference: &inference}, nil
}

func (qs queryServer) GetInferencesAtBlock(
	ctx context.Context,
	req *types.QueryInferencesAtBlockRequest,
) (*types.QueryInferencesAtBlockResponse, error) {

	inferences, err := qs.k.GetInferencesAtBlock(ctx, req.TopicId, req.BlockHeight)
	if err != nil {
		return nil, err
	}

	return &types.QueryInferencesAtBlockResponse{Inferences: inferences}, nil
}

// Return full set of inferences in I_i from the chain
func (qs queryServer) GetNetworkInferencesAtBlock(
	ctx context.Context,
	req *types.QueryNetworkInferencesAtBlockRequest,
) (*types.QueryNetworkInferencesAtBlockResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	topic, err := qs.k.GetTopic(ctx, req.TopicId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "topic %v not found", req.TopicId)
	}
	if topic.EpochLastEnded == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"network inference not available for topic %v",
			req.TopicId,
		)
	}
	if req.BlockHeightLastInference > sdkCtx.BlockHeight() {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"block height cannot be greater than current block height %v",
			sdkCtx.BlockHeight(),
		)
	}

	networkInferences, err := synth.GetNetworkInferencesAtBlock(
		sdk.UnwrapSDKContext(ctx),
		qs.k,
		req.TopicId,
		req.BlockHeightLastInference,
		req.BlockHeightLastReward,
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryNetworkInferencesAtBlockResponse{NetworkInferences: networkInferences}, nil
}
