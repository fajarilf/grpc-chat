export interface ApiResponse<TModel> {
    status: string,
    data: TModel
}

export interface ApiCursorResponse<TModel> extends ApiResponse<TModel> {
    nextCursor: number
    prevCursor: number
    hasMore: boolean
}