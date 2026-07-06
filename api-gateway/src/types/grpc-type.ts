import { status as GrpcStatus } from '@grpc/grpc-js';

export interface Message {
    username: string,
    content: string,
    timestamp: number,
    room: string
}

export interface MessageResponse {
    success: boolean,
    message_id: string
}

export interface JoinRequest {
    username: string,
    room: string
}

export interface HistoryRequest {
    room: string;
    limit: number;
}

export interface HistoryResponse {
    message: Message[]
}

/**
 * Converts a gRPC status code to the equivalent HTTP status code.
 * Based on the standard mapping used by grpc-gateway / Google API guidelines.
 * https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
 */
export function grpcToHttpStatus(grpcCode: number): number {
    switch (grpcCode) {
        case GrpcStatus.OK:
            return 200;
        case GrpcStatus.CANCELLED:
            return 499; // Client Closed Request (nginx convention)
        case GrpcStatus.UNKNOWN:
            return 500;
        case GrpcStatus.INVALID_ARGUMENT:
            return 400;
        case GrpcStatus.DEADLINE_EXCEEDED:
            return 504;
        case GrpcStatus.NOT_FOUND:
            return 404;
        case GrpcStatus.ALREADY_EXISTS:
            return 409;
        case GrpcStatus.PERMISSION_DENIED:
            return 403;
        case GrpcStatus.UNAUTHENTICATED:
            return 401;
        case GrpcStatus.RESOURCE_EXHAUSTED:
            return 429;
        case GrpcStatus.FAILED_PRECONDITION:
            return 400;
        case GrpcStatus.ABORTED:
            return 409;
        case GrpcStatus.OUT_OF_RANGE:
            return 400;
        case GrpcStatus.UNIMPLEMENTED:
            return 501;
        case GrpcStatus.INTERNAL:
            return 500;
        case GrpcStatus.UNAVAILABLE:
            return 503;
        case GrpcStatus.DATA_LOSS:
            return 500;
        default:
            return 500;
    }
}