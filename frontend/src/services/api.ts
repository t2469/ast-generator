export interface ASTNode {
    type: string
    start_byte: number;
    end_byte: number;
    children?: ASTNode[];
}

export interface ParseRequest {
    language: string;
    code: string;
}

export async function parseCode(request: ParseRequest): Promise<ASTNode> {
    const response = await fetch("http://localhost:8080/parse", {
        method: "POST",
        headers: {"Content-Type": "application/json",},
        body: JSON.stringify(request),
    });
    if (!response.ok) {
        throw new Error(`parse API Error: ${response.status}`);
    }
    return response.json();
}