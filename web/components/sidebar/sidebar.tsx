import { Box } from '../styles/box';
import { Sidebar } from './sidebar.styles';
import { Flex } from '../styles/flex';
import { CompaniesDropdown } from './companies-dropdown';
import { HomeIcon } from '../icons/sidebar/home-icon';
import { ProductsIcon } from '../icons/sidebar/products-icon';
import { ReportsIcon } from '../icons/sidebar/reports-icon';
import { SettingsIcon } from '../icons/sidebar/settings-icon';
import { SidebarItem } from './sidebar-item';
import { SidebarMenu } from './sidebar-menu';
import { useSidebarContext } from '../layout/layout-context';
import { useRouter } from 'next/router';
import { DarkModeSwitch } from '../navbar_dashboard/darkmodeswitch';

export const SidebarWrapper = () => {
   const router = useRouter();
   const { collapsed, setCollapsed } = useSidebarContext();

   return (
      <Box
         as="aside"
         css={{
            height: '100vh',
            zIndex: 202,
            position: 'sticky',
            top: '0',
         }}
      >
         {collapsed ? <Sidebar.Overlay onClick={setCollapsed} /> : null}

         <Sidebar collapsed={collapsed}>
            <Sidebar.Header>
               <CompaniesDropdown />
            </Sidebar.Header>
            <Flex
               direction={'column'}
               justify={'between'}
               css={{ height: '100%' }}
            >
               <Sidebar.Body className="body sidebar">
                  <SidebarItem
                     title="Home"
                     icon={<HomeIcon />}
                     isActive={router.pathname === '/dashboard'}
                     href="/dashboard"
                  />
                  <SidebarMenu title="Model">
                     <SidebarItem
                        isActive={router.pathname === '/submit'}
                        title="Submit Model"
                        icon={<ProductsIcon />}
                        href="submit"
                     />
                     <SidebarItem
                        isActive={router.pathname === '/list'}
                        title="List Model"
                        icon={<ReportsIcon />}
                        href="list"
                     />
                  </SidebarMenu>

                  <SidebarMenu title="User">
                     <SidebarItem
                        isActive={router.pathname === '/settings'}
                        title="Settings"
                        icon={<SettingsIcon />}
                        href="#"
                     />
                  </SidebarMenu>

               </Sidebar.Body>
               <Sidebar.Footer>
                  <SidebarItem
                     title="Settings"
                     icon={<SettingsIcon />}
                     isActive={router.pathname === '/settings'}
                     href="#"
                  />
                  <DarkModeSwitch />
               </Sidebar.Footer>
            </Flex>
         </Sidebar>
      </Box>
   );
};
