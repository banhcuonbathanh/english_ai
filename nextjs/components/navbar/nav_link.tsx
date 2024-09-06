"use client";

import React from "react";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { cn } from "@/lib/utils";
import { existingRoutes } from "@/lib/navbar/nav_link_data";

const NavLinks = () => {
  const pathname = usePathname();
  const router = useRouter();

  const handleLinkClick = (href: string, label: string) => {
    router.push(`${href}`);
  };

  return (
    <>
      {existingRoutes.map((route) => (
        <div key={route.href} className="flex flex-col items-center gap-x-6">
          <Link
            href={route.href}
            className={cn(
              "text-sm font-medium transition-colors hover:text-primary",
              pathname === route.href
                ? "border-b-2 border-orange-500 transition-all ease-in-out duration-300 "
                : "transition-all ease-in-out duration-300 text-muted-foreground"
            )}
            onClick={() => handleLinkClick(route.href, route.label)}
            style={{ cursor: "pointer" }}
          >
            {route.label}
          </Link>
        </div>
      ))}
    </>
  );
};

export default NavLinks;