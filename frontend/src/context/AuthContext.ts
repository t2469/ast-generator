import { createContext } from 'react';
import { UserInfo } from '../services/api';

interface AuthContextType {
    user: UserInfo | null;
    setUser: React.Dispatch<React.SetStateAction<UserInfo | null>>;
}

export const AuthContext = createContext<AuthContextType>({
    user: null,
    setUser: () => { },
}); 