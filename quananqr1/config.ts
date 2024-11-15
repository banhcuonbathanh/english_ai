import { z } from "zod";

const configSchema = z.object({
  NEXT_PUBLIC_API_ENDPOINT: z.string(),
  NEXT_PUBLIC_URL: z.string(),
  NEXT_PUBLIC_API_Create_User: z.string(),
  NEXT_PUBLIC_API_Get_Account_Email: z.string(),

  NEXT_PUBLIC_API_Login: z.string(),
  NEXT_PUBLIC_API_Logout: z.string(),
  NEXT_PUBLIC_Image_Upload: z.string(),
  NEXT_PUBLIC_Add_Dished: z.string(),

  NEXT_PUBLIC_Add_Guest_login: z.string(),

  NEXT_PUBLIC_Get_Dished_intenal: z.string(),
  NEXT_PUBLIC_Upload: z.string(),
  NEXT_PUBLIC_Folder1_BE: z.string(),

  NEXT_PUBLIC_Table_List: z.string(),

  NEXT_PUBLIC_intern_table_end_point: z.string(),

  NEXT_PUBLIC_Table_End_Point: z.string(),
  NEXT_PUBLIC_Set_End_Point: z.string(),

  NEXT_PUBLIC_Get_set_intenal: z.string(),

  Order_Internal_End_Point: z.string(),
  Order_External_End_Point: z.string(),
  NEXT_PUBLIC_API_Guest_Login: z.string(),
  NEXT_PUBLIC_API_Guest_Logout: z.string(),
  wslink: z.string(),
  wsAuth: z.string()
});
// /users/email
const configProject = configSchema.safeParse({
  wslink: "ws://localhost:8888/ws/",
  wsAuth: "ws/api/ws-auth", 
  Order_Internal_End_Point: "api/order",
  Order_External_End_Point: "orders",
  NEXT_PUBLIC_API_Guest_Login: "qr/guest/login",
  NEXT_PUBLIC_API_Guest_Logout: "qr/guest/logout",

  NEXT_PUBLIC_Get_set_intenal: process.env.NEXT_PUBLIC_Get_set_intenal,

  NEXT_PUBLIC_Set_End_Point: process.env.NEXT_PUBLIC_Set_End_Point,
  NEXT_PUBLIC_Table_End_Point: process.env.NEXT_PUBLIC_Table_End_Point,

  NEXT_PUBLIC_intern_table_end_point:
    process.env.NEXT_PUBLIC_intern_table_end_point,
  NEXT_PUBLIC_Table_List: process.env.NEXT_PUBLIC_Table_List,
  NEXT_PUBLIC_Folder1_BE: process.env.NEXT_PUBLIC_Folder1_BE,
  NEXT_PUBLIC_Upload: process.env.NEXT_PUBLIC_Upload,
  NEXT_PUBLIC_API_ENDPOINT: process.env.NEXT_PUBLIC_API_ENDPOINT,
  NEXT_PUBLIC_URL: process.env.NEXT_PUBLIC_URL,
  NEXT_PUBLIC_API_Create_User: process.env.NEXT_PUBLIC_API_Create_User,

  NEXT_PUBLIC_API_Get_Account_Email:
    process.env.NEXT_PUBLIC_API_Get_Account_Email,
  NEXT_PUBLIC_API_Login: process.env.NEXT_PUBLIC_API_Login,

  NEXT_PUBLIC_API_Logout: process.env.NEXT_PUBLIC_API_Logout,
  NEXT_PUBLIC_Image_Upload: process.env.NEXT_PUBLIC_Image_Upload,

  NEXT_PUBLIC_Add_Dished: process.env.NEXT_PUBLIC_Add_Dished,

  NEXT_PUBLIC_Add_Guest_login: process.env.NEXT_PUBLIC_Add_Guest_login,
  NEXT_PUBLIC_Get_Dished_intenal: process.env.NEXT_PUBLIC_Get_Dished_intenal
});

if (!configProject.success) {
  console.error(configProject.error.errors);
  throw new Error("Các khai báo biến môi trường không hợp lệ");
}

const customConfig = {
  NEXT_PUBLIC_Get_set_intenal: process.env.NEXT_PUBLIC_Get_set_intenal
};

const envConfig = configProject.data;

export default envConfig;
