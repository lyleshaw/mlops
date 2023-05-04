// @ts-nocheck

import { Avatar, Dropdown, Navbar, Text } from '@nextui-org/react';
import { useState, useEffect } from 'react';
import router from 'next/router';

export const UserDropdown = () => {
   const [user, setUser] = useState(null);

   useEffect(() => {
      const localUser = localStorage.getItem('user');
      if (localUser) {
         setUser(JSON.parse(localUser));
      } else {
         router.push('/') 
      }
   }, []);
   return (
      <Dropdown placement="bottom-right">
         <Navbar.Item>
            <Dropdown.Trigger>
               <Avatar
                  bordered
                  as="button"
                  color="secondary"
                  size="md"
                  src={user?.avatar || 'https://avatars.githubusercontent.com/u/3333'}
               />
            </Dropdown.Trigger>
         </Navbar.Item>
         <Dropdown.Menu
            aria-label="User menu actions"
            onAction={(actionKey) => console.log({ actionKey })}
         >
            <Dropdown.Item key="profile" css={{ height: '$18' }}>
               <Text b color="inherit" css={{ d: 'flex' }}>
                  Signed in as
               </Text>
               <Text b color="inherit" css={{ d: 'flex' }}>
               {user?.username || 'Nil'}
               </Text>
            </Dropdown.Item>
            <Dropdown.Item key="settings" withDivider>
               My Settings
            </Dropdown.Item>
            <Dropdown.Item key="help_and_feedback" withDivider>
               Help & Feedback
            </Dropdown.Item>
            <Dropdown.Item key="logout" withDivider color="error">
               Log Out
            </Dropdown.Item>
         </Dropdown.Menu>
      </Dropdown>
   );
};
