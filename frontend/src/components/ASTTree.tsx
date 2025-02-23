import React, { useEffect, useRef, useState } from "react";
import Tree from "react-d3-tree";
import { ASTNode } from "../services/api";

interface ASTTreeProps {
    node: ASTNode;
}

interface TreeNode {
    name: string;
    attributes?: { [key: string]: string | number | boolean };
    children?: TreeNode[];
}

const transformASTToTree = (node: ASTNode): TreeNode => {
    return {
        name: node.type,
        attributes: {
            content: node.content,
        },
        children: node.children ? node.children.map(transformASTToTree) : [],
    };
};

interface CustomNodeElementProps {
    nodeDatum: TreeNode;
}

//文字数が10文字以上かどうかで変更予定
const CustomNodeElement: React.FC<CustomNodeElementProps> = ({ nodeDatum }) => {
    return (
        <g>
            <rect
                x="-50"
                y="-45"
                rx="5"
                ry="5"
                style={{
                    width: "100",
                    height: '90',
                    fill: 'turquoise',
                    stroke: 'salmon',
                    strokeWidth: '3'
                }}
            />
            <foreignObject x="-50" y="-45" width="100px" height="90px">
                <div
                    className="w-[100px] h-[90px] flex flex-col items-center justify-center break-words text-center"
                    style={{ wordBreak: 'break-word' }}
                >
                    <span className="font-bold">{nodeDatum.name}</span>
                    {nodeDatum.attributes && nodeDatum.attributes.content && (
                        <span className="text-sm">{nodeDatum.attributes.content}</span>
                    )}
                </div>
            </foreignObject>
        </g>
    )
}

const ASTTree: React.FC<ASTTreeProps> = ({ node }) => {
    const [translate, setTranslate] = useState({ x: 0, y: 0 });
    const treeContainer = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (treeContainer.current) {
            const dimensions = treeContainer.current.getBoundingClientRect();
            setTranslate({
                x: dimensions.width / 2,
                y: dimensions.height / 2,
            });
        }
    }, []);

    const treeData = transformASTToTree(node);

    return (
        <div id="treeWrapper" style={{ width: "100%", height: "500px", border: "1px solid #ccc", backgroundColor: "#bbb" }} ref={treeContainer}>
            <Tree
                data={treeData}
                renderCustomNodeElement={({ nodeDatum }) =>
                    <CustomNodeElement nodeDatum={nodeDatum} />
                }
                translate={translate}
                orientation="vertical"
                pathFunc="step"
            />
        </div>
    );
};

export default ASTTree;
