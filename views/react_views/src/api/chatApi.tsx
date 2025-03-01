export type ChatMessage = {
    chat_id: number;
    chat_detail: string;
    author_id: number;
};

export const fetchChatMessage = async (): Promise<ChatMessage[]> => {
    const response = await fetch("http://localhost:8080/api/chat_history");
    if (!response.ok) {
        throw new Error("データの取得に失敗しました");
      }
    return response.json();
} 

export const postChatMessage = async (message: string): Promise<ChatMessage> => {
    if (!message.trim()) {
        throw new Error("メッセージは空文字で送信できません");
    }

    const response = await fetch("http://localhost:8080/api/chat_post/",{
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({chat_detail: message}),        
    });

    if (!response.ok) {
        throw new Error("送信に失敗しました");
      }
    return response.json();
};

export const updateChatMessage = async (chatId: number, message: string): Promise<ChatMessage> => {
    if (!message.trim()) {
        throw new Error("メッセージは空文字で送信できません");
    }

    const response = await fetch(`http://localhost:8080/api/chat_edit/${chatId}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({chat_detail: message}),
    });

    if (!response.ok) {
        throw new Error("更新に失敗しました");
    }
    return response.json();
}

export const deleteChatMessage = async (chatId: number): Promise<number> => {
    const response = await fetch (`http://localhost:8080/api/chat_delete/${chatId}`, { method: "DELETE"})

    if (!response.ok) {
        throw new Error("削除に失敗しました");
    }
    return response.json();
};