import { useRef, useEffect } from "react";

const LogoutModal = (props: any) => {
    const modalRef = useRef<HTMLDivElement>(null);

    const updateModalPosition = () => {
        if (modalRef.current && props.targetRef.current) {
            const iconRect = props.targetRef.current.getBoundingClientRect();
            modalRef.current.style.top = `${iconRect.bottom}px`;
            modalRef.current.style.left = `${iconRect.left-45}px`;
        }
    };

    useEffect(() => {
        if (props.showFlag) {
            updateModalPosition();
            window.addEventListener('resize', updateModalPosition);
            window.addEventListener('scroll', updateModalPosition);
        }

        return () => {
            updateModalPosition();
            window.removeEventListener('resize', updateModalPosition);
            window.removeEventListener('scroll', updateModalPosition);
        };
    }, [props.showFlag]);

    return (
        <>
            {props.showFlag ? (
                <div className="fixed inset-0 flex justify-center items-start z-50">
                    <div className="bg-white p-5 rounded shadow-lg fixed z-60" ref={modalRef}>
                        <button
                            onClick={props.handleLogout}
                            className="text-gray-700 hover:text-red-500 transition-colors cursor-pointer mb-2"
                        >ログアウト</button>
                        <br/>
                        <button
                            onClick={() => props.setShowModal(false)}
                            className="text-gray-700 hover:text-blue-500 transition-colors cursor-pointer"
                        >Close</button>
                    </div>
                </div>
            ) : null}
        </>
    );
};

export default LogoutModal;