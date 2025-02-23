import React, { useEffect, useRef, useState } from "react";
import Tree from "react-d3-tree";
import { ASTNode } from "../services/api";

interface TreeNode {
    name: string;
    attributes?: { [key: string]: string | number | boolean };
    children?: TreeNode[];
}

const transformASTToTree = (node: ASTNode): TreeNode => {
    return {
        name: node.type,
        attributes: {
            id: node.id,
            content: node.content,
        },
        children: node.children ? node.children.map(transformASTToTree) : [],
    };
};

const traverseASTPreOrder = (node: ASTNode, order: number[] = []): number[] => {
    order.push(node.id);
    if (node.children) {
        node.children.forEach(child => traverseASTPreOrder(child, order));
    }
    return order;
};

interface CustomNodeElementProps {
    nodeDatum: TreeNode;
    highlighted: boolean;
}

const CustomNodeElement: React.FC<CustomNodeElementProps> = ({ nodeDatum, highlighted }) => {
    return (
        <g>
            <circle
                r="32"
                fill={highlighted ? "#60A5FA" : "#ffffff"}
                stroke="black"
                strokeWidth="2"
            />
            <text
                x="0"
                y="-40"
                textAnchor="middle"
                fontSize="18"
                fill="#333"
                style={{
                    whiteSpace: "pre-line",
                    fontWeight: 300,
                    fontFamily: "sans-serif",
                    stroke: "none",
                }}
            >
                {nodeDatum.name}
            </text>
            <text
                x="0"
                y="50"
                textAnchor="middle"
                fontSize="16"
                fill="#333"
                style={{
                    whiteSpace: "pre-line",
                    fontWeight: 300,
                    fontFamily: "sans-serif",
                    stroke: "none",
                }}
            >
                {nodeDatum.attributes?.content}
            </text>
        </g>
    );
};

interface ASTTreeProps {
    node: ASTNode;
}

const ASTTree: React.FC<ASTTreeProps> = ({ node }) => {
    const [translate, setTranslate] = useState({ x: 0, y: 0 });
    const [currentHighlightId, setCurrentHighlightId] = useState<number | null>(null);
    const [dfsOrder, setDfsOrder] = useState<number[]>([]);
    const treeContainer = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (treeContainer.current) {
            const dimensions = treeContainer.current.getBoundingClientRect();
            setTranslate({
                x: dimensions.width / 2,
                y: dimensions.height / 8,
            });
        }
    }, []);

    useEffect(() => {
        const order = traverseASTPreOrder(node);
        setDfsOrder(order);
    }, [node]);

    useEffect(() => {
        if (dfsOrder.length === 0) return;
        let index = 0;
        setCurrentHighlightId(dfsOrder[index]);
        const interval = setInterval(() => {
            index++;
            if (index >= dfsOrder.length) {
                clearInterval(interval);
            } else {
                setCurrentHighlightId(dfsOrder[index]);
            }
        }, 1000);
        return () => clearInterval(interval);
    }, [dfsOrder]);

    const treeData = transformASTToTree(node);

    const containerStyles = {
        width: "100%",
        height: "800px",
        border: "1px solid #ccc",
        backgroundColor: "#ffffff",
    };

    return (
        <div id="treeWrapper" style={containerStyles} ref={treeContainer}>
            <Tree
                data={treeData}
                translate={translate}
                orientation="vertical"
                pathFunc="elbow"
                collapsible
                transitionDuration={1250}
                enableLegacyTransitions
                renderCustomNodeElement={({ nodeDatum }) => (
                    <CustomNodeElement
                        nodeDatum={nodeDatum}
                        highlighted={nodeDatum.attributes?.id === currentHighlightId}
                    />
                )}
                zoomable
            />
        </div>
    );
};

export default ASTTree;
