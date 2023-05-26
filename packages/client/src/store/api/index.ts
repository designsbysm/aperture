import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

import env from '../../environment';
import { Appointment } from '../../models';

const appointments: Appointment[] = [
  {
    date: '2021-10-12T09:15:00-07:00',
    id: 10341,
    lab: {
      name: 'Embry Onsite',
    },
    patients: [
      {
        name: 'Doe, Jane',
        result: 'cancelled',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test',
            paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia',
            paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John',
        result: 'cancelled',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test',
            paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia',
            paidBy: 'patient',
          },
        ],
      },
    ],
    site: {
      address: '1833 W. Southern Ave.',
      city: 'Mesa',
      name: 'Mesa Community College',
      notes: 'Enter from Southern Ave. adjacent to football field.',
      state: 'AZ',
      zip: '85205',
    },
  },
  {
    date: '2021-11-14T16:45:00-07:00',
    id: 10672,
    lab: {
      name: 'Embry Onsite',
    },
    patients: [
      {
        name: 'Doe, John',
        result: 'negative',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test',
            paidBy: 'insurance',
          },
        ],
      },
    ],
    site: {
      address: '1833 W. Southern Ave.',
      city: 'Mesa',
      name: 'Mesa Community College',
      notes: 'Enter from Southern Ave. adjacent to football field.',
      state: 'AZ',
      zip: '85205',
    },
  },
  {
    date: '2021-12-28T10:00:00-07:00',
    id: 10833,
    lab: {
      name: 'Health Track Rx',
    },
    patients: [
      {
        name: 'Doe, Jane',
        result: 'negative',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
        ],
      },
      {
        name: 'Doe, John',
        result: 'negative',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
        ],
      },
      {
        name: 'Doe, John Jr.',
        result: 'cancelled',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
        ],
      },
      {
        name: 'Doe, Jill',
        result: 'cancelled',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
        ],
      },
    ],
    site: {
      address: '1833 W. Southern Ave.',
      city: 'Mesa',
      name: 'Mesa Community College',
      state: 'AZ',
      zip: '85205',
    },
  },
  {
    date: '2022-02-14T07:30:00-07:00',
    id: 11494,
    lab: {
      name: 'Health Track Rx',
    },
    patients: [
      {
        name: 'Doe, Jane',
        result: 'positive',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John',
        result: 'positive',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John Jr.',
        result: 'positive',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, Jill',
        result: 'positive',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
    ],
    site: {
      address: '16875 Canyon Trails Blvd.',
      city: 'Goodyear',
      name: 'Copper Trails School',
      state: 'AZ',
      zip: '85338',
    },
  },
  {
    date: '2022-03-01T15:30:00-07:00',
    id: 13935,
    lab: {
      name: 'Health Track Rx',
    },
    patients: [
      {
        name: 'Doe, Jane',
        result: 'negative',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John',
        result: 'negative',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John Jr.',
        result: 'positive',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, Jill',
        result: 'positive',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
    ],
    site: {
      address: '16875 Canyon Trails Blvd.',
      city: 'Goodyear',
      name: 'Copper Trails School',
      state: 'AZ',
      zip: '85338',
    },
  },
  {
    date: '2022-03-12T12:45:00-07:00',
    id: 14316,
    lab: {
      name: 'Embry Lab',
    },
    patients: [
      {
        name: 'Doe, John Jr.',
        result: 'positive',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
    ],
    site: {
      address: '3825 N. 67th Ave.',
      city: 'Phoenix',
      name: 'Watts Family Maryvale YMCA',
      state: 'AZ',
      zip: '85033',
    },
  },
  {
    date: '2022-03-14T09:50:00-07:00',
    id: 14747,
    lab: {
      name: 'Phoenix Toxicology & Lab Services',
    },
    patients: [
      {
        name: 'Doe, Jane',
        result: 'negative',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
    ],
    site: {
      address: '3000 N Dysart Rd.',
      city: 'Avondale',
      name: 'Estrella Mountain Community College',
      notes: 'North side, behind the police station.',
      state: 'AZ',
      zip: '85392',
    },
  },
  {
    date: '2022-04-07T14:45:00-07:00',
    id: 17188,
    lab: {
      name: 'Sanora Quest Laboratories',
    },
    patients: [
      {
        name: 'Doe, Jane',
        result: 'negative',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'free',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John',
        result: 'negative',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'free',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John Jr.',
        result: 'positive',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'free',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, Jill',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'free',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
    ],
    site: {
      address: '3900 S. 55th Ave.',
      city: 'Phoenix',
      name: 'Maricopa Institute of Technology',
      state: 'AZ',
      zip: '85043',
    },
  },
  {
    date: '2022-04-14T15:30:00-07:00',
    id: 18649,
    lab: {
      name: 'Phoenix Toxicology & Lab Services',
    },
    patients: [
      {
        name: 'Doe, Jane',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
        ],
      },
      {
        name: 'Doe, John',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
        ],
      },
      {
        name: 'Doe, John Jr.',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
        ],
      },
      {
        name: 'Doe, Jill',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'insurance',
          },
        ],
      },
    ],
    site: {
      address: '1122 S 67th Ave.',
      city: 'Phoenix',
      name: 'South Ridge High School',
      state: 'AZ',
      zip: '85043',
    },
  },
  {
    date: '2023-04-15T11:00:00-07:00',
    id: 21350,
    lab: {
      name: 'Sanora Quest Laboratories',
    },
    patients: [
      {
        name: 'Doe, Jane',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'free',
          },
          {
            cost: 200,
            name: 'Covid-19 Rapid PCR Test Accula', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'free',
          },
          {
            cost: 200,
            name: 'Covid-19 Rapid PCR Test Accula', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, John Jr.',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'free',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
      {
        name: 'Doe, Jill',
        result: 'pending',
        tests: [
          {
            cost: 0,
            name: 'COVID-19 PCR Diagnostic Test', paidBy: 'free',
          },
          {
            cost: 100,
            name: 'COVID-19 Rapid Antigen Test Sofia', paidBy: 'patient',
          },
        ],
      },
    ],
    site: {
      address: '1833 W. Southern Ave.',
      city: 'Mesa',
      name: 'Mesa Community College',
      notes: 'Enter from Southern Ave. adjacent to football field.',
      state: 'AZ',
      zip: '85205',
    },
  },
];

export const apiSlice = createApi({
  baseQuery: fetchBaseQuery({
    baseUrl: `${env.apiURL}`,
  }),
  endpoints: builder => ({
    getAppointment: builder.query({
      query: id => `/v1/appointments/${id}`,
      transformResponse: res => appointments.find(appointment => appointment.id === res),
    }),
    getAppointments: builder.query({
      query: () => '/v1/appointments',
      transformResponse: () => appointments,
    }),
  }),
  reducerPath: 'api',
});

export const { useGetAppointmentQuery, useGetAppointmentsQuery } = apiSlice;
