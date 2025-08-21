import { IUser } from '../../../entities/user';

export interface IUserWithPassword extends IUser {
	password?: string;
}
