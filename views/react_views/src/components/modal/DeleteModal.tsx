import Modal from "react-modal";

type DeleteModalProps = {
    deleteId: number | null;
    setDeleteId: (deleteId: number | null) => void;
    handleConfirmedDelete: (chatId: number) => void;
};

Modal.setAppElement("#root");

export default function DeleteModal({ deleteId, setDeleteId, handleConfirmedDelete }: DeleteModalProps) {
    return (
        <Modal
            isOpen={deleteId !== null}
            onRequestClose={() => setDeleteId(null)}
            className="modal"
            overlayClassName="overlay"
        >
            <div className="modal-header">
                <h2>削除しますか？</h2>
            </div>
            <button onClick={() => deleteId !== null && handleConfirmedDelete(deleteId)}>確定</button>
            <button onClick={() => setDeleteId(null)}>キャンセル</button>
        </Modal>
    );
}