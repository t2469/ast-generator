import React from "react";
import { ASTNode } from "../services/api";

interface ASTTreeProps {
    node: ASTNode;
}

const ASTTree: React.FC<ASTTreeProps> = ({ node }) => {
    return (
        <div style={{ marginLeft: "20px", borderLeft: "1px solid #ccc", paddingLeft: "10px" }}>
            <div>
                <strong>{node.type}</strong> ({node.start_byte} - {node.end_byte})
            </div>
            {node.children && node.children.length > 0 && (
                <div>
                    {node.children.map((child, index) => (
                        <ASTTree key={index} node={child} />
                    ))}
                </div>
            )}
        </div>
    );
};

export default ASTTree;
