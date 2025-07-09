"use client";

import { Moon, Sun } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useTheme } from "@/components/theme-provider";

export function ModeToggle() {
  const { setTheme } = useTheme();
  const theme = useTheme().theme;

  return (
    <Button
      size={"icon"}
      onClick={() => (theme == "dark" ? setTheme("light") : setTheme("dark"))}
      className="cursor-pointer"
    >
      <Sun className="absolute h-[1.2rem] w-[1.2rem] scale-100 transition-all dark:scale-0" />
      <Moon className="absolute h-[1.2rem] w-[1.2rem] scale-0 transition-all dark:scale-100" />
    </Button>
  );
}
