export interface IUser {
	username: string;
	firstName?: string;
	lastName?: string;
	birthday?: Date;
	gender?: string;
	interests?: string[];
	city?: string;
}

export interface IUserWithPassword extends IUser {
	password?: string;
}
