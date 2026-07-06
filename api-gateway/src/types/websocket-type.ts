export interface Message {
    user_id: number,
    message: string,
    timestamp: Date
}

export interface Incoming {
    type: "message" | "reply"
    room_id: number
    message: Message
}