import { postChatMessage } from "../api/chatApi";
import { useForm } from 'react-hook-form';
import "./../App.css";

type Props = {
    onNewMessage: (message: string) => void;
};

type FormData = {
    message: string;
}

export default function ChatForm( { onNewMessage }: Props) {
    const { register, handleSubmit, reset } = useForm<FormData>();

    const onSubmit = async (data: FormData) => {
        if (!data.message.trim()) return;

        try {
            const newChat = await postChatMessage(data.message);
            console.log("送信したデータ:", newChat);
            onNewMessage(newChat.chat_detail);
            reset(); // フォームをリセット
        } catch (error) {
            console.error("送信エラー:", error);
        }
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="chatter_form">
          <textarea {...register('message', { required: true })} />
          <input type="submit" value="post!" />
        </form>
      );
}