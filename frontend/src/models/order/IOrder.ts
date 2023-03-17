import { UsersInterface } from "../user/IUser";

export interface OrdersInterface {
    ID: number,
    User_ID: number,
    User: UsersInterface,
}