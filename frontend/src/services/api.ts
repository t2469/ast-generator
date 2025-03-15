export interface ASTNode {
    id: number;
    type: string;
    start_byte: number;
    end_byte: number;
    content: string;
    children?: ASTNode[];
}

export interface ParseRequest {
    language: string;
    code: string;
}

export interface UserInfo {
    name: string;
    email: string;
    picture: string;
    sub: string;
}

export interface SourceCode {
    id: number;
    title: string;
    description: string;
    language: string;
    code: string;
    created_at: string;
}

const BASE_URL = import.meta.env.VITE_API_URL

export async function parseCode(request: ParseRequest): Promise<ASTNode> {
    const response = await fetch(`${BASE_URL}/parse`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(request),
    });
    if (!response.ok) {
        throw new Error(`parse API Error: ${response.status}`);
    }
    return response.json() as Promise<ASTNode>;
}

export async function getCurrentUser(): Promise<UserInfo> {
    try {
        const response = await fetch(`${BASE_URL}/auth/current_user`, {
            method: "GET",
            credentials: "include",
        });

        if (!response.ok) {
            throw new Error(`get current user API Error: ${response.status}`);
        }
        return response.json() as Promise<UserInfo>;
    } catch (error) {
        console.error("Error fetching user info:", error);
        throw error;
    }
}

export async function logout(): Promise<boolean> {
    try {
        const response = await fetch(`${BASE_URL}/auth/logout`, {
            method: "GET",
            credentials: "include",
        });
        return response.ok;
    } catch (error) {
        console.error("Logout failed:", error);
        return false;
    }
}

export async function saveSourceCode(data: {
    title: string;
    description: string;
    language: string;
    code: string;
}): Promise<SourceCode> {
    const response = await fetch(`${BASE_URL}/source_codes/save`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        credentials: "include",
        body: JSON.stringify(data),
    });

    if (!response.ok) {
        throw new Error("Failed to save source code");
    }
    return response.json();
}

export async function deleteSourceCode(id: number): Promise<void> {
    const response = await fetch(`${BASE_URL}/source_codes/${id}`, {
        method: "DELETE",
        credentials: "include",
    });

    if (!response.ok) {
        throw new Error("Failed to delete source code");
    }
}

export async function getUserSourceCodes(): Promise<SourceCode[]> {
    const response = await fetch(`${BASE_URL}/source_codes/user`, {
        method: "GET",
        credentials: "include",
    });
    if (!response.ok) {
        throw new Error("Failed to fetch source codes");
    }
    return response.json();
}