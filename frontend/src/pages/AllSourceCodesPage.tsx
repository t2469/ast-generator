import React, { useEffect, useState } from "react";
import { SourceCode, getAllSourceCodes } from "../services/api";

const AllSourceCodesPage: React.FC = () => {
    const [sourceCodes, setSourceCodes] = useState<SourceCode[]>([]);
    const [error, setError] = useState<string>("");

    useEffect(() => {
        getAllSourceCodes()
            .then((data) => setSourceCodes(data))
            .catch((err) => setError(err instanceof Error ? err.message : "Error fetching source codes"));
    }, []);

    return (
        <div className="container mx-auto p-6 pt-20">
            <h2 className="text-2xl font-bold mb-4">保存されたプログラム一覧</h2>
            {error && <p className="text-red-500">{error}</p>}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {sourceCodes.map((code) => (
                    <div key={code.id} className="bg-white p-4 rounded shadow">
                        <h3 className="text-xl font-semibold">{code.title}</h3>
                        <p className="text-gray-600">{code.description}</p>
                        <pre className="bg-gray-100 p-2 mt-2 rounded overflow-x-auto">
                            {code.code}
                        </pre>
                        <p className="text-sm text-gray-500 mt-1">
                            {new Date(code.created_at).toLocaleString()}
                        </p>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default AllSourceCodesPage;
