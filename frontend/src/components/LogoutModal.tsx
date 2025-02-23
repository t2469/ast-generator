import { useRef, useEffect } from "react";

const LogoutModal = (props: any) => {
    const modalRef = useRef<HTMLDivElement>(null);

    const updateModalPosition = () => {
        if (modalRef.current && props.targetRef.current) {
            const iconRect = props.targetRef.current.getBoundingClientRect();
            modalRef.current.style.top = `${iconRect.bottom+15}px`;
            modalRef.current.style.left = `${iconRect.left - 45}px`;
        }
    };

    const handleClickOutside = (event: MouseEvent) => {
        if (modalRef.current && !modalRef.current.contains(event.target as Node)) {
            props.setShowModal(false);
        }
    };

    useEffect(() => {
        if (props.showFlag) {
            updateModalPosition();
            window.addEventListener('resize', updateModalPosition);
            window.addEventListener('scroll', updateModalPosition);
            document.addEventListener('mousedown', handleClickOutside);
        }

        return () => {
            window.removeEventListener('resize', updateModalPosition);
            window.removeEventListener('scroll', updateModalPosition);
            document.removeEventListener('mousedown', handleClickOutside);
        };
    }, [props.showFlag]);

    return (
        <>
            {props.showFlag ? (
                <div className="fixed inset-0 flex justify-center items-start z-50">
                    <div className="bg-white p-5 rounded shadow-lg fixed z-60" ref={modalRef}>
                        <button
                            onClick={props.handleLogout}
                            className="text-gray-700 hover:text-red-500 transition-colors cursor-pointer m-0"
                        >ログアウト</button>
                    </div>
                </div>
            ) : null}
        </>
    );
};

export default LogoutModal;