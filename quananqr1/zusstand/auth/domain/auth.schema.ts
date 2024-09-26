import { RoleValues } from "@/constants/type";
import z from "zod";

export const LoginBody = z
  .object({
    email: z.string().email(),
    password: z.string().min(6).max(100)
  })
  .strict();

export type LoginBodyType = z.TypeOf<typeof LoginBody>;

export const LoginRes = z.object({
  data: z.object({
    accessToken: z.string(),
    refreshToken: z.string(),
    account: z.object({
      id: z.number(),
      name: z.string(),
      email: z.string(),
      role: z.enum(RoleValues),
      avatar: z.string().nullable()
    })
  }),
  message: z.string()
});

export type LoginResType = z.TypeOf<typeof LoginRes>;

export const RefreshTokenBody = z
  .object({
    refreshToken: z.string()
  })
  .strict();

export type RefreshTokenBodyType = z.TypeOf<typeof RefreshTokenBody>;

export const RefreshTokenRes = z.object({
  data: z.object({
    accessToken: z.string(),
    refreshToken: z.string()
  }),
  message: z.string()
});

export type RefreshTokenResType = z.TypeOf<typeof RefreshTokenRes>;

export const LogoutBody = z
  .object({
    refreshToken: z.string()
  })
  .strict();

export type LogoutBodyType = z.TypeOf<typeof LogoutBody>;

export const RegisterBody = z.object({
  name: z.string().min(1, "Name is required"),
  email: z.string().email("Invalid email address"),
  password: z.string().min(8, "Password must be at least 8 characters long"),
  is_admin: z.boolean(),
  phone: z.number(),
  image: z.string().optional(),
  address: z.string().min(1, "Address is required"),
  created_at: z.string().datetime(),
  updated_at: z.string().datetime()
});

// Derive the type from the schema
export type RegisterBodyType = z.infer<typeof RegisterBody>;
