import React, { useState, useContext } from "react";
import { ASTNode, parseCode, saveSourceCode } from "../services/api";
import ASTTree from "../components/ASTTree";
import { AuthContext } from "../context/AuthContext";

const ASTPage: React.FC = () => {
    const [code, setCode] = useState<string>("");
    const [language, setLanguage] = useState<string>("go");
    const [ast, setAst] = useState<ASTNode | null>(null);
    const [error, setError] = useState<string>("");
    const [saveMessage, setSaveMessage] = useState<string>("");

    const [showSaveModal, setShowSaveModal] = useState<boolean>(false);
    const [modalTitle, setModalTitle] = useState<string>("");
    const [modalDescription, setModalDescription] = useState<string>("");
    const [modalLanguage, setModalLanguage] = useState<string>(language);
    const [modalCode, setModalCode] = useState<string>(code);

    const { user } = useContext(AuthContext);

    const languageOptions = [
        { value: "bash", label: "Bash" },
        { value: "c", label: "C" },
        { value: "cpp", label: "C++" },
        { value: "css", label: "CSS" },
        { value: "dockerfile", label: "Dockerfile" },
        { value: "go", label: "Go" },
        { value: "html", label: "HTML" },
        { value: "java", label: "Java" },
        { value: "javascript", label: "JavaScript" },
        { value: "kotlin", label: "Kotlin" },
        { value: "php", label: "PHP" },
        { value: "python", label: "Python" },
        { value: "ruby", label: "Ruby" },
        { value: "rust", label: "Rust" },
        { value: "sql", label: "SQL" },
        { value: "yaml", label: "YAML" },
    ];

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("");
        setSaveMessage("");
        try {
            const result = await parseCode({ language, code });
            setAst(result);
        } catch (err: unknown) {
            setError(err instanceof Error ? err.message : "不明なエラーが発生しました");
        }
    };

    const openSaveModal = () => {
        setModalTitle("");
        setModalDescription("");
        setModalLanguage(language);
        setModalCode(code);
        setShowSaveModal(true);
    };

    const handleModalSave = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("");
        try {
            await saveSourceCode({
                title: modalTitle,
                description: modalDescription,
                language: modalLanguage,
                code: modalCode,
            });
            setSaveMessage("プログラムの保存に成功しました！");
            setShowSaveModal(false);
        } catch (err: unknown) {
            setError(err instanceof Error ? err.message : "保存中に不明なエラーが発生しました");
        }
    };

    return (
        <div className="relative">
            {!showSaveModal && (
                <div className="container mx-auto p-6 pt-20">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                        {/* コード入力フォーム */}
                        <div className="bg-white p-6 rounded-lg shadow">
                            <h2 className="text-2xl font-bold mb-4">コード入力</h2>
                            <form onSubmit={handleSubmit} className="space-y-4">
                                <div>
                                    <label className="block text-sm font-medium text-gray-700 mb-1">
                                        言語:
                                    </label>
                                    <select
                                        value={language}
                                        onChange={(e) => {
                                            setLanguage(e.target.value);
                                            setModalLanguage(e.target.value);
                                        }}
                                        className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                                    >
                                        {languageOptions.map((option) => (
                                            <option key={option.value} value={option.value}>
                                                {option.label}
                                            </option>
                                        ))}
                                    </select>
                                </div>
                                <div>
                                    <label className="block text-sm font-medium text-gray-700 mb-1">
                                        コード:
                                    </label>
                                    <textarea
                                        value={code}
                                        onChange={(e) => {
                                            setCode(e.target.value);
                                            setModalCode(e.target.value);
                                        }}
                                        rows={10}
                                        className="block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                                        placeholder="ここにコードを入力してください..."
                                    />
                                </div>
                                <div className="flex space-x-4">
                                    <button
                                        type="submit"
                                        className="flex-1 py-2 font-semibold text-white rounded-md shadow bg-gradient-to-r from-blue-500 to-indigo-600 hover:from-blue-600 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-200"
                                    >
                                        解析する
                                    </button>
                                    {user && (
                                        <button
                                            type="button"
                                            onClick={openSaveModal}
                                            className="flex-1 py-2 font-semibold text-white rounded-md shadow bg-gradient-to-r from-green-500 to-teal-600 hover:from-green-600 hover:to-teal-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-all duration-200"
                                        >
                                            保存する
                                        </button>
                                    )}
                                </div>
                                {error && (
                                    <p className="text-red-500 text-sm">エラー: {error}</p>
                                )}
                                {saveMessage && (
                                    <p className="text-green-500 text-sm">{saveMessage}</p>
                                )}
                            </form>
                        </div>

                        {/* AST 表示 */}
                        <div className="bg-white p-6 rounded-lg shadow">
                            <h2 className="text-2xl font-bold mb-4">構文木 (AST) 表示</h2>
                            <div className="p-4 border border-gray-300 rounded-md min-h-[300px] bg-white">
                                {ast ? (
                                    <ASTTree node={ast} />
                                ) : (
                                    <p className="text-gray-500">解析結果がここに表示されます</p>
                                )}
                            </div>
                        </div>
                    </div>
                </div>
            )}
            {showSaveModal && (
                <div className="fixed inset-0 z-50 flex items-center justify-center">
                    <div className="absolute inset-0 bg-black opacity-50"></div>
                    <div className="relative bg-white rounded-lg shadow-lg p-8 w-full max-w-3xl min-h-[400px]">
                        <h2 className="text-xl font-bold mb-4">プログラム保存</h2>
                        <form onSubmit={handleModalSave} className="space-y-4">
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                    タイトル:
                                </label>
                                <input
                                    type="text"
                                    value={modalTitle}
                                    onChange={(e) => setModalTitle(e.target.value)}
                                    className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                                    placeholder="プログラムのタイトル"
                                    required
                                />
                            </div>
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                    説明:
                                </label>
                                <textarea
                                    value={modalDescription}
                                    onChange={(e) => setModalDescription(e.target.value)}
                                    rows={3}
                                    className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                                    placeholder="プログラムの説明"
                                />
                            </div>
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                    言語:
                                </label>
                                <select
                                    value={modalLanguage}
                                    onChange={(e) => setModalLanguage(e.target.value)}
                                    className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                                >
                                    <option value="go">Go</option>
                                    <option value="cpp">C++</option>
                                </select>
                            </div>
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                    コード:
                                </label>
                                <textarea
                                    value={modalCode}
                                    onChange={(e) => setModalCode(e.target.value)}
                                    rows={8}
                                    className="block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                                />
                            </div>
                            <div className="flex justify-end space-x-4">
                                <button
                                    type="button"
                                    onClick={() => setShowSaveModal(false)}
                                    className="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400 transition-colors"
                                >
                                    キャンセル
                                </button>
                                <button
                                    type="submit"
                                    className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600 transition-colors"
                                >
                                    保存
                                </button>
                            </div>
                        </form>
                        {error && (
                            <p className="text-red-500 text-sm mt-2">エラー: {error}</p>
                        )}
                    </div>
                </div>
            )}
        </div>
    );
};

export default ASTPage;
