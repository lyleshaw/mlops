// @ts-nocheck

import { Table } from '@nextui-org/react';
import { Box } from '../styles/box';
import { columns } from './data';
import { RenderCell } from './render-cell';
import { getAppAPI } from '../../api/my-app';
import { useEffect, useState } from 'react';
import router from 'next/router';

export const TableWrapper = () => {
   const [user, setUser] = useState(null);
   const [models, setModels] = useState([]);
   const [reloading, setReloading] = useState(false);

   useEffect(() => {
      const localUser = localStorage.getItem('user');
      if (localUser) {
         setUser(JSON.parse(localUser));
      } else {
         router.push('/')
      }
      const fetchApp = async () => {
         setReloading(false);
         let response = await getAppAPI();
         console.log('response', response)
         const temp = response.data.map((item) => {
            return {
               id: item.app_id,
               name: item.app_name,
               status: item.app_status,
               port: item.app_port,
               image: item.app_image,
               user: item.user_id,
               avatar: user ? user.avatar : 'https://avatars.githubusercontent.com/u/25427168?v=4',
               create: item.create_at,
               update: item.update_at,
            }
         })
         setModels(temp);
      }
      fetchApp();
   }, [reloading])

   return (
      <Box
         css={{
            '& .nextui-table-container': {
               boxShadow: 'none',
            },
         }}
      >
         <Table
            aria-label="Example table with custom cells"
            css={{
               height: 'auto',
               minWidth: '100%',
               boxShadow: 'none',
               width: '100%',
               px: 0,
            }}
            color='secondary'
            selectionMode="multiple"
         >
            <Table.Header columns={columns}>
               {(column) => (
                  <Table.Column
                     key={column.uid}
                     hideHeader={column.uid === 'actions'}
                     align={column.uid === 'actions' ? 'center' : 'start'}
                  >
                     {column.name}
                  </Table.Column>
               )}
            </Table.Header>
            <Table.Body items={models}>
               {(item) => (
                  <Table.Row>
                     {(columnKey) => (
                        <Table.Cell>
                           {RenderCell({ setReloading: setReloading, model: item, columnKey: columnKey })}
                        </Table.Cell>
                     )}
                  </Table.Row>
               )}
            </Table.Body>
            <Table.Pagination
               shadow
               noMargin
               align="center"
               rowsPerPage={8}
               onPageChange={(page) => console.log({ page })}
            />
         </Table>
      </Box>
   );
};
