import { Button, Input, Text, Grid, Spacer, Link, Card } from '@nextui-org/react';
import { HouseIcon } from '../icons/breadcrumb/house-icon';
import { UsersIcon } from '../icons/breadcrumb/users-icon';
import { Flex } from '../styles/flex';
import { TableWrapper } from '../table/table';

export const Lists = () => {
   const handler = () => window.location.href = '/submit';
   return (
      <Flex
         css={{
            'mt': '$5',
            'px': '$6',
            '@sm': {
               mt: '$10',
               px: '$16',
            },
         }}
         justify={'center'}
         direction={'column'}
      >
         <Grid.Container gap={0} justify="flex-start">
            <Grid xs={1}>
               <HouseIcon />
               <Link href="/dashboard" color="text">Home</Link>
               <Text>&nbsp;/</Text>
            </Grid>
            <Grid xs={1}>
               <UsersIcon />
               <Link href="#" color="text">Model</Link>
               <Text>&nbsp;/</Text>
            </Grid>
            <Grid xs={1}>
               <Link href="#" color="text">List</Link>
            </Grid>
         </Grid.Container>

         <Card>
            <Card.Body>
               <Text h3>List Models</Text>
               <Grid.Container gap={2} justify="center">
                  <Grid xs={4} css={{ justifyContent: "flex-start" }}>
                     <Input
                        css={{ width: '100%', maxW: '500px' }}
                        placeholder="Search Models"
                     />
                  </Grid>
                  <Grid xs={6} css={{ justifyContent: "center" }}>
                  </Grid>
                  <Grid xs={2} css={{ justifyContent: "flex-end" }}>
                     <Button auto color="gradient" onPress={handler}>
                        Add Model
                     </Button>
                  </Grid>
               </Grid.Container>
               <Grid css={{ justifyContent: "center" }}>
                  <TableWrapper />
               </Grid>
               <Spacer y={3}></Spacer>
            </Card.Body>
         </Card>
      </Flex>
   );
};
