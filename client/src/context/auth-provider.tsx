import React, { createContext, useContext, useEffect, useState } from "react";

type AuthProviderProps = {
  children: React.ReactNode;
  defaultAuth?: boolean;
  storageKey?: string;
};

type AuthProvider = {
  isAuth: boolean;
  setAuth: (authContext: string | null) => void;
};

const initialAuthState: AuthProvider = {
  isAuth: false,
  setAuth: () => null,
};

// eslint-disable-next-line react-refresh/only-export-components
export const AuthContext = createContext<AuthProvider>(initialAuthState);

export const AuthProvider = ({
  children,
  defaultAuth = false,
  storageKey = "access-token",
}: AuthProviderProps) => {
  const [isAuth, setAuth] = useState<boolean>(
    localStorage.getItem(storageKey) != null ? true : defaultAuth
  );

  useEffect(() => {
    const token = localStorage.getItem(storageKey);

    setAuth(!!token);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const value = {
    isAuth,
    setAuth: (token: string | null) => {
      if (token === null) {
        localStorage.removeItem(storageKey);
        setAuth(false);
        return;
      }

      localStorage.setItem(storageKey, token);
      setAuth(true);
    },
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

// eslint-disable-next-line react-refresh/only-export-components
export const useAuth = () => {
  const context = useContext(AuthContext);

  if (context === undefined)
    throw new Error("useTheme must be used within a AuthProvider");

  return context;
};
