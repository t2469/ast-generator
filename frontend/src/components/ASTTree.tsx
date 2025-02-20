import React, {useEffect, useRef, useState} from "react";
import Tree from "react-d3-tree";
import {ASTNode} from "../services/api";

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
        children: node.children ? node.children.map(transformASTToTree) : [],
    };
};

const ASTTree: React.FC<ASTTreeProps> = ({node}) => {
    const [translate, setTranslate] = useState({x: 0, y: 0});
    const treeContainer = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (treeContainer.current) {
            const dimensions = treeContainer.current.getBoundingClientRect();
            setTranslate({
                x: dimensions.width,
                y: dimensions.height,
            });
        }
    }, []);

    const treeData = transformASTToTree(node);

    return (
        <div id="treeWrapper" style={{width: "100%", height: "500px", border: "1px solid #ccc"}} ref={treeContainer}>
            <Tree
                data={treeData}
                translate={translate}
                orientation="vertical"
                pathFunc="elbow"
            />
        </div>
    );
};

export default ASTTree;
