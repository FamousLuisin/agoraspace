import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "@/App.tsx";
import { ThemeProvider } from "./components/theme-provider.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <div className="w-full min-h-screen h-full flex flex-col items-center bg-background dark:bg-background">
        <App />
      </div>
    </ThemeProvider>
  </StrictMode>
);
