// @ts-nocheck

import { ChangeEvent, useState } from 'react';
import {
   FormElement,
   Modal,
   Input,
   Button,
   Text,
   Spacer,
   Grid
} from '@nextui-org/react';
import { toast } from 'react-toastify';
import { IconButton } from '../table/table.styled';
import { EditIcon } from '../icons/table/edit-icon';
import { updateAppAPI } from '../../api/update-app';

interface Props {
   setReloading: any,
   modelID: any;
}

export const UpdateModel = ({setReloading, modelID}: Props) => {
   console.log('props', modelID)
   const [visible, setVisible] = useState(false);
   const [values, setValues] = useState({
      appName: '',
      dockerImage: '',
   })
   const handler = () => setVisible(true);
   const closeHandler = () => {
      setVisible(false);
   }
   const handleChange = (prop: string) => (event: ChangeEvent<FormElement>) => {
      setValues({ ...values, [prop]: event.target.value })
   }

   const updateApp = async () => {
      try {
         const response = await updateAppAPI(modelID, values.appName, values.dockerImage, false);
         console.log(response);
         if (response.message === 'success') {
            setReloading(true);
            setVisible(false);
         }

      } catch (error) {
         console.log(error);
         toast.error('失败，请重试！');
      }
   };

   return (
      <div>
         <IconButton
            onClick={() => handler()}
         >
            <EditIcon size={20} fill="#979797" />
         </IconButton>
         <Modal
            closeButton
            blur
            aria-labelledby="modal-title"
            open={visible}
            onClose={closeHandler}
         >
            <Modal.Header>
               <Text
                  h1
                  transform='uppercase'
                  id="modal-title"
                  size={32}
                  css={{
                     textGradient: "45deg, $blue600 10%, $pink600 70%",
                  }}>
                  Update Model
               </Text>
            </Modal.Header>
            <Modal.Body>
               <Grid.Container gap={1} justify="flex-start">
                  <Grid xs={2}>
                  </Grid>
                  <Grid xs={8} css={{ justifyContent: "center" }}>
                     <Input
                        underlined
                        labelLeft="Model Name"
                        placeholder="GPT4"
                        width="100%"
                        color="secondary"
                        onChange={handleChange('appName')}
                        value={values.appName}
                     />
                  </Grid>
                  <Grid xs={2}>
                  </Grid>
                  <Grid xs={2}>
                  </Grid>
                  <Grid xs={8} css={{ justifyContent: "center" }}>
                     <Input
                        underlined
                        labelLeft="Docker Image"
                        placeholder="lyle/gradio:latest"
                        width="100%"
                        color="secondary"
                        onChange={handleChange('dockerImage')}
                        value={values.dockerImage}
                     />
                  </Grid>
                  <Grid xs={2}>
                  </Grid>
               </Grid.Container>
               <Spacer y={3}></Spacer>
            </Modal.Body>
            <Modal.Footer>
               <Button auto flat bordered color='gradient' onPress={closeHandler}>
                  Close
               </Button>
               <Button auto shadow color='gradient' onPress={updateApp}>
                  Submit
               </Button>
            </Modal.Footer>
         </Modal>
      </div>
   );
};
