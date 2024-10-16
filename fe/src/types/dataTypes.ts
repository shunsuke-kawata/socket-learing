// src/types.ts
export interface StatusNoAt {
  ID: number;
  StatusName: string;
}

export interface AddressNoAt {
  ID: number;
  IPAddress: string;
  ColorCode: string;
}

export interface TaskNoAt {
  ID: number;
  Title: string;
  Description: string;
  StatusID: number;
  Status: StatusNoAt;
  AddressID: number;
  Address: AddressNoAt;
}
