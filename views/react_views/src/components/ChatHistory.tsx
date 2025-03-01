import { useEffect, useState } from "react";
import { ChatMessage, fetchChatMessage, deleteChatMessage, updateChatMessage } from "../api/chatApi";
import ChatMessageItem from "./ChatMessage";
import EditModal from "./modal/EditModal";
import DeleteModal from "./modal/DeleteModal";
import "./../App.css";

type Props = {
    refresh: boolean;
    setRefresh: (refresh: boolean) => void;
};

export default function ChatHistory({ refresh, setRefresh }: Props) {
    const [chatMessages, setChatMessages] = useState<ChatMessage[]>([]);
    const [editId, setEditId] = useState<number | null>(null);
    const [deleteId, setDeleteId] = useState<number | null>(null);
    const [editText, setEditText] = useState("");

    useEffect(() => {
        const fetchChatHistory = async () => {
            try {
                console.log("Fetching chat history...");
                const messages = await fetchChatMessage();
                console.log("取得したデータ:", messages);
                setChatMessages(messages);
            } catch (error) {
                console.error("チャット履歴の取得エラー:", error);
            }
        };
        fetchChatHistory();
    }, [refresh]);

    const handleDelete = (deleteId: number) => {
        setDeleteId(deleteId);
    };
    const handleConfirmedDelete = async (chatId: number) => {
        try {
            const id = await deleteChatMessage(chatId);
            console.log("削除しました:", id);
            setDeleteId(null);
            setRefresh(!refresh);
        } catch (error) {
            console.error("削除エラー:", error);
        }
    };

    const handleEdit = (chatId: number, detail: string) => {
        setEditId(chatId);
        setEditText(detail);
    };
    const handleUpdate = async () => {
        try {
            if (editId !== null) {
                const edit_chat = await updateChatMessage(editId, editText);
                console.log("更新しました:", edit_chat);
            }
            setEditId(null);
            setRefresh(!refresh);
        } catch (error) {
            console.error("更新エラー:", error);
        }
    };

    return (
        <div className="chat-box">
            {chatMessages.map((msg) => (
                <ChatMessageItem key={msg.chat_id} message={msg} onDelete={handleDelete} onEdit={handleEdit} />
            ))}
            <EditModal
                editId={editId}
                editText={editText}
                setEditId={setEditId}
                setEditText={setEditText}
                handleUpdate={handleUpdate}
            />
            <DeleteModal
                deleteId={deleteId}
                setDeleteId={setDeleteId}
                handleConfirmedDelete={handleConfirmedDelete}
            />
        </div>
    );
}
