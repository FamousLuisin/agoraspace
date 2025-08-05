"use client";

import { Button } from "@/components/ui/button";
import { ModeToggle } from "@/components/mode-toggle";
import { Link } from "react-router-dom";
import { useAuth } from "@/context/auth-provider";

export default function Header() {
  const { isAuth, setAuth } = useAuth();

  const logoutAuth = () => {
    setAuth(null);
  };

  return (
    <header className="w-11/12 flex justify-between py-4 border-b-1 border-zinc-500 items-center">
      <div className="flex items-center">
        <Link to="/">
          <img
            src="src/assets/light_column.png"
            alt=""
            className="w-10 h-10 block dark:hidden"
          />
        </Link>
        <Link to="/">
          <img
            src="src/assets/dark_column.png"
            alt=""
            className="w-10 h-10 hidden dark:block"
          />
        </Link>
        <Link to="/">
          <h1 className="text-primary dark:text-primary text-3xl font-serif font-semibold px-2">
            Ágora
          </h1>
        </Link>
      </div>
      <div className="flex gap-2">
        {!isAuth && (
          <nav className="flex gap-2">
            <Button asChild variant="default">
              <Link to="/signin">Sign in</Link>
            </Button>
            <Button asChild variant="secondary">
              <Link to="/signup">Sign up</Link>
            </Button>
          </nav>
        )}
        {isAuth && (
          <nav className="flex gap-2">
            <Button asChild variant="secondary" onClick={logoutAuth}>
              <Link to="/">Logout</Link>
            </Button>
          </nav>
        )}
        <ModeToggle />
      </div>
    </header>
  );
}
