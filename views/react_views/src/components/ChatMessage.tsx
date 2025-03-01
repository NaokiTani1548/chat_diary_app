import { ChatMessage } from "../api/chatApi";
import "./../App.css";

type Props = {
    message: ChatMessage;
    onDelete: (chatId: number) => void;
    onEdit: (chatId: number, detail: string) => void;
};

export default function ChatMessageItem({ message, onDelete, onEdit }: Props) {
    return (
        <div className="message" >
            <div className="avatar"></div>
            <div className="text">
                <div className="chat_text">
                    {message.chat_detail}
                </div>
                <div className="button_container">
                    <button className="chat_edit_button" onClick={() => onEdit(message.chat_id, message.chat_detail)}>
                        ✎
                    </button>
                    <button className="chat_delete_button" onClick={() => onDelete(message.chat_id)}>
                        -
                    </button>
                </div>
            </div>
        </div>
    );
}