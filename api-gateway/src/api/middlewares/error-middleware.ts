import { status as GrpcStatus, ServiceError } from '@grpc/grpc-js';
import { NextFunction, Request, Response } from "express";
import { grpcToHttpStatus } from '../../types/grpc-type';

function isGrpcError(error: unknown): error is ServiceError {
    return (
        error instanceof Error &&
        'code' in error &&
        Object.values(GrpcStatus).includes((error as ServiceError).code)
    );
}

export async function ErrorMiddleware(error: Error, req: Request, res: Response, next: NextFunction) {
    if (isGrpcError(error)) {
        const httpStatus = grpcToHttpStatus(error.code)
        return res.status(httpStatus).json({
            code: httpStatus,
            status: "error",
            error: error.details || 'Unknow error',
        })
    }

    res.status(500).json({
        code: 500,
        status: "error",
        error: 'Internal server error'
    })
}