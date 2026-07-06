export interface ApiResponse<TModel> {
    status: string,
    data: TModel
    paging?: Paging
}

export interface Paging {
    nextCursor: number
    prevCursor: number
    hasMore: boolean
}