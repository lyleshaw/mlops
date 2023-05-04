import { Button, Divider, Text } from '@nextui-org/react';
import React from 'react';
import { BoxIcon } from '../icons/BoxIcon';
import { Box } from '../styles/box';
import { Flex } from '../styles/flex';

export const Faq = () => {
   return (
      <>
         <Flex
            css={{
               py: '$20',
               gap: '$18',
               px: '$6',
            }}
            direction={'column'}
         >
            <Flex align={'center'} direction={'column'}>
               <Text span css={{ color: '$blue600' }}>
                  FAQ
               </Text>
               <Text h2>You Have Questions?</Text>
               <Text
                  span
                  css={{
                     maxWidth: '700px',
                     color: '$accents8',
                     textAlign: 'center',
                  }}
               >
                  Do not Worrier.
               </Text>
            </Flex>

            <Flex
               css={{
                  'gap': '$10',
                  '@lg': {
                     px: '$64',
                  },
               }}
               direction={'column'}
            >
               <Flex css={{ gap: '$5' }} justify={'center'}>
                  <BoxIcon />
                  <Flex direction={'column'} css={{ gap: '$3' }}>
                     <Text h3>
                        Q: What is MLOps?
                     </Text>
                     <Text
                        span
                        css={{
                           color: '$accents8',
                        }}
                     >
                        A: MLOps, or Machine Learning Operations, is the practice of automating and optimizing the end-to-end machine learning lifecycle. It aims to reduce the friction of moving from model development to production and enable continuous delivery of machine learning.
                     </Text>
                  </Flex>
               </Flex>

               <Flex css={{ gap: '$5' }} justify={'center'}>
                  <BoxIcon />
                  <Flex direction={'column'} css={{ gap: '$3' }}>
                     <Text h3>
                        Q: Do I need to be a machine learning expert to use MLOps?
                     </Text>
                     <Text
                        span
                        css={{
                           color: '$accents8',
                        }}
                     >
                        A: No, MLOps aims to make machine learning more accessible to everyone. While knowledge of ML concepts is helpful, MLOps platforms abstract away much of the complexity involved in building and deploying ML models. You can focus on your area of expertise and let MLOps handle the ML pipelines.
                     </Text>
                  </Flex>
               </Flex>

               <Flex css={{ gap: '$5' }} justify={'center'}>
                  <BoxIcon />
                  <Flex direction={'column'} css={{ gap: '$3' }}>
                     <Text h3>
                        Q: What are the benefits of MLOps?
                     </Text>
                     <Text
                        span
                        css={{
                           color: '$accents8',
                        }}
                     >
                        A: Key benefits of MLOps include:
                        • Accelerate time to market for ML models
                        • Increase productivity of data scientists
                        • Improve collaboration between data scientists and operations teams
                        • Ensure governance, auditability and compliance of ML models
                        • Monitor ML models in production for performance drift and retrain as needed
                        • Scale ML deployments across the organization
                     </Text>
                  </Flex>
               </Flex>
            </Flex>
         </Flex>

         <Divider
            css={{ position: 'absolute', inset: '0p', left: '0', mt: '$5' }}
         />
      </>
   );
};
