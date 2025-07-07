import { Copyright } from "lucide-react";

export default function Footer() {
  return (
    <footer className="w-full flex py-4 justify-center gap-2 items-center border-t-1 border-zinc-500">
      <span className="hidden dark:inline">
        <Copyright color="#f7f7f7" />
      </span>
      <span className="inline dark:hidden">
        <Copyright color="#343434" />
      </span>
      <span className="text-primary">Noki</span>
    </footer>
  );
}
