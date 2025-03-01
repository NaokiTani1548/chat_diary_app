import Modal from "react-modal";

type EditModalProps = {
    editId: number | null;
    editText: string;
    setEditId: (editId: number | null) => void;
    setEditText: (editText: string) => void;
    handleUpdate: () => void;
};

Modal.setAppElement("#root");

export default function EditModal({ editId, editText, setEditId, setEditText, handleUpdate }: EditModalProps) {
    return (
        <Modal
            isOpen={editId !== null}
            onRequestClose={() => setEditId(null)}
            className="modal"
            overlayClassName="overlay"
        >
            <div className="modal-header">
                <h2>メッセージ編集</h2>
            </div>
            <textarea value={editText} onChange={(e) => setEditText(e.target.value)} />
            <button onClick={handleUpdate}>更新</button>
            <button onClick={() => setEditId(null)}>キャンセル</button>
        </Modal>
    );
}