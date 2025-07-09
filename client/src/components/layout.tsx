"use client";

import Header from "@/components/header";
import Footer from "@/components/footer";
import { Outlet } from "react-router-dom";

export default function Layout() {
  return (
    <div className="w-full min-h-screen h-full flex flex-col items-center bg-background dark:bg-background">
      <Header />
      <Outlet />
      <Footer />
    </div>
  );
}
