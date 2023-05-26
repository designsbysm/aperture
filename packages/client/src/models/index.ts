export type Appointment = {
  id: number;
  date: string;
  lab: {
    name: string;
  },
  patients: Patient[],
  site: {
    address: string;
    city: string;
    name: string;
    notes?: string;
    state: string;
    zip: string;
  };
}

type Patient = {
  name: string;
  result: string;
  tests: Test[];
}

type Test = {
  name: string;
  cost: number;
  paidBy: string;
}
