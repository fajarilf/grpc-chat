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