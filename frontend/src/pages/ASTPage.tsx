import React, {useState} from "react";
import {ASTNode, parseCode} from "../services/api";
import ASTTree from "../components/ASTTree";

const ASTPage: React.FC = () => {
    const [code, setCode] = useState<string>("");
    const [language, setLanguage] = useState<string>("go");
    const [ast, setAst] = useState<ASTNode | null>(null);
    const [error, setError] = useState<string>("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("");
        try {
            const result = await parseCode({language, code});
            setAst(result);
        } catch (err: unknown) {
            setError(err instanceof Error ? err.message : '不明なエラーが発生しました');
        }
    };

    return (
        <div style={{display: "flex", padding: "20px"}}>
            <div style={{flex: 1, marginRight: "20px"}}>
                <h2>コード入力</h2>
                <form onSubmit={handleSubmit}>
                    <div>
                        <label>言語:</label>
                        <select value={language} onChange={(e) => setLanguage(e.target.value)}>
                            <option value="go">Go</option>
                            <option value="cpp">C++</option>
                        </select>
                    </div>
                    <div>
                        <label>コード:</label>
                        <textarea
                            value={code}
                            onChange={(e) => setCode(e.target.value)}
                            rows={10}
                            cols={50}
                        />
                    </div>
                    <button type="submit">解析する</button>
                </form>
                {error && <p style={{color: "red"}}>エラー: {error}</p>}
            </div>
            <div style={{flex: 1}}>
                <h2>構文木(AST)表示</h2>
                {ast ? <ASTTree node={ast}/> : <p>解析結果がここに表示されます</p>}
            </div>
        </div>
    );
};

export default ASTPage;
