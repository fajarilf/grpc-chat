export type SuccessResponse<T> = {
    status: string;
    data: T;
    paging?: Paging;
}

export type Paging = {
    nextCursor: number;
    prevCursor: number;
    hasMore: boolean;
}

export type ListFilter = {
    forward: boolean;
    cursor: number;
    size: number
}