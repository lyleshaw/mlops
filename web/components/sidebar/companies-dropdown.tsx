import { Text } from '@nextui-org/react';
import { MLLogo } from '../icons/mllogo';
import { Box } from '../styles/box';
import { Flex } from '../styles/flex';

export const CompaniesDropdown = () => {
   return (
            <Box>
               <Flex align={'center'} css={{ gap: '$7' }}>
                  <MLLogo />
                  <Box>
                     <Text
                        h3
                        size={'$xl'}
                        weight={'medium'}
                        css={{
                           m: 0,
                           color: '$accents9',
                           lineHeight: '$lg',
                           mb: '-$5',
                        }}
                     >
                        MLOps
                     </Text>
                     <Text
                        span
                        weight={'medium'}
                        size={'$xs'}
                        css={{ color: '$accents8' }}
                     >
                        Welcome To Use!
                     </Text>
                  </Box>
               </Flex>
            </Box>
   );
};
