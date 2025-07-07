import { Button } from "@/components/ui/button";
import { ModeToggle } from "@/components/mode-toggle";

export default function Header() {
  return (
    <header className="w-full flex justify-between py-4 border-b-1 border-zinc-500 items-center">
      <div className="flex items-center gap-2">
        <img
          src="src/assets/light_column.png"
          alt=""
          className="w-10 h-10 block dark:hidden"
        />
        <img
          src="src/assets/dark_column.png"
          alt=""
          className="w-10 h-10 hidden dark:block"
        />
        <h1 className="text-primary dark:text-primary text-3xl font-serif font-semibold">
          Ágora
        </h1>
      </div>
      <div className="flex">
        <nav className="flex gap-2">
          <Button asChild variant="default">
            <a href="/login">Sign in</a>
          </Button>
          <Button asChild variant="secondary">
            <a href="/login">Sign up</a>
          </Button>
          <ModeToggle />
        </nav>
      </div>
    </header>
  );
}
