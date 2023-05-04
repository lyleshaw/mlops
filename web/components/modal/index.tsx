// @ts-nocheck

import React, { ChangeEvent } from 'react';
import {
   FormElement,
   Modal,
   Input,
   Row,
   Checkbox,
   Button,
   Text,
   Navbar,
   Spacer,
   Col
} from '@nextui-org/react';
import { loginAPI } from '../../api/login';
import { registerAPI } from '../../api/register';
import { sendVerificationCodeAPI } from '../../api/email-code';
import router from 'next/router';
import { toast } from 'react-toastify';
import { getUserAPI } from '../../api/get-user';

export const ModalLogin = () => {
   const [visible, setVisible] = React.useState(false);
   const [values, setValues] = React.useState({
      email: '',
      password: '',
   })
   const handler = () => setVisible(true);
   const closeHandler = () => {
      setVisible(false);
   };
   const handleChange = (prop: string) => (event: ChangeEvent<FormElement>) => {
      setValues({ ...values, [prop]: event.target.value })
   }
   const login = async () => {
      try {
         const response = await loginAPI(values.email, values.password);
         console.log(response);
         if (response.message === 'success') {
            localStorage.setItem('token', response.data.token);
            localStorage.setItem('expire', response.data.expire);
            const userResp = await getUserAPI();
            localStorage.setItem('user', JSON.stringify(userResp.data))
            router.push('/dashboard');
         }
      } catch (error) {
         console.log(error);
         toast.error('登录失败，请重试！');
      }
   };
   return (
      <div>
         <Navbar.Link block onPress={handler}>
            Login
         </Navbar.Link>
         <Modal
            closeButton
            blur
            aria-labelledby="modal-title"
            open={visible}
            onClose={closeHandler}
         >
            <Modal.Header>
               <Text
                  transform='uppercase'
                  id="modal-title"
                  size={18}
                  css={{
                     textGradient: "45deg, $blue600 10%, $pink600 70%",
                  }}>
                  Welcome to &nbsp;
                  <Text
                     transform='uppercase'
                     b
                     size={18}
                     css={{
                        textGradient: "45deg, $blue600 -10%, $pink600 70%",
                     }}>
                     MLOps System
                  </Text>
               </Text>
            </Modal.Header>
            <Modal.Body>
               <Spacer y={0.5} />
               <Input
                  autoFocus
                  fullWidth
                  id='email'
                  underlined
                  labelPlaceholder="Email"
                  color="secondary"
                  onChange={handleChange('email')}
                  value={values.email}
               />
               <Spacer y={0.5} />
               <Input.Password
                  fullWidth
                  id='password'
                  underlined
                  labelPlaceholder="Password"
                  color="secondary"
                  onChange={handleChange('password')}
                  value={values.password}
               />
               <Row justify="space-between">
                  <Checkbox color='gradient'>
                     <Text transform='uppercase' size={14}>Remember me</Text>
                  </Checkbox>
                  <Text transform='uppercase' size={14}>Forgot password?</Text>
               </Row>
            </Modal.Body>
            <Modal.Footer>
               <Button auto flat bordered color='gradient' onPress={closeHandler}>
                  Close
               </Button>
               <Button auto shadow color='gradient' onPress={login}>
                  Sign in
               </Button>
            </Modal.Footer>
         </Modal>
      </div>
   );
};

export const ModalRegister = () => {
   const [visible, setVisible] = React.useState(false);
   const [values, setValues] = React.useState({
      username: '',
      email: '',
      code: '',
      password: '',
   })
   const handler = () => setVisible(true);
   const closeHandler = () => {
      setVisible(false);
   };
   const handleChange = (prop: string) => (event: ChangeEvent<FormElement>) => {
      setValues({ ...values, [prop]: event.target.value })
   }
   const sendVerificationCode = async () => {
      await sendVerificationCodeAPI(values.email)
   }
   const register = async () => {
      try {
         const response = await registerAPI(values.username, values.email, values.password, values.code)
         console.log(response);
         if (response.message === 'success') {
            const loginResp = await loginAPI(values.email, values.password);
            localStorage.setItem('token', loginResp.data.token);
            localStorage.setItem('expire', loginResp.data.expire);
            localStorage.setItem('user', JSON.stringify(response.data))
            router.push('/dashboard');
         }
      } catch (error) {
         console.log(error);
         toast.error('登录失败，请重试！');
      }
   }
   return (
      <div>
         <Button auto flat onPress={handler} color={"gradient"}
            css={{
               color: "$white",
            }}>
            Register
         </Button>
         <Modal
            closeButton
            blur
            aria-labelledby="modal-title"
            open={visible}
            onClose={closeHandler}
         >
            <Modal.Header>
               <Text
                  transform='uppercase'
                  b
                  size={18}
                  css={{
                     textGradient: "45deg, $blue600 -10%, $pink600 70%",
                  }}>
                  Start Your MLOps Trip HERE🚀
               </Text>
            </Modal.Header>
            <Modal.Body>
               <Spacer y={0.5} />
               <Input
                  autoFocus
                  fullWidth
                  id='username'
                  underlined
                  labelPlaceholder="User Name"
                  color="secondary"
                  onChange={handleChange('username')}
                  value={values.username}
               />
               <Spacer y={0.5} />
               <Input.Password
                  fullWidth
                  id='password'
                  underlined
                  labelPlaceholder="Password"
                  color="secondary"
                  onChange={handleChange('password')}
                  value={values.password}
               />
               <Spacer y={0.5} />
               <Input
                  fullWidth
                  id='email'
                  underlined
                  labelPlaceholder="Email"
                  color="secondary"
                  onChange={handleChange('email')}
                  value={values.email}
               />
               <Spacer y={0.5} />
               <Row justify="center" align="center">
                  <Col span={8}>
                     <Input
                        fullWidth
                        id='code'
                        maxLength={6}
                        underlined
                        labelPlaceholder="Verification Code"
                        color="secondary"
                        onChange={handleChange('code')}
                        value={values.code}
                     />
                  </Col>
                  <Col span={5}>
                     <div style={{ display: "flex", justifyContent: "flex-end" }}>
                        <Button
                           color="gradient"
                           bordered
                           auto
                           onPress={sendVerificationCode}
                        >
                           <Text transform='uppercase' css={{ textGradient: "45deg, $pink600 -20%, $blue600 100%", }}>Send Code</Text>
                        </Button>
                     </div>
                  </Col>
               </Row>
               <Row justify="space-between">
                  <Text transform='uppercase' b size={14}>Already have an account? Sign in instead</Text>
               </Row>
            </Modal.Body>
            <Modal.Footer>
               <Button auto flat bordered color='gradient' onPress={closeHandler}>
                  Close
               </Button>
               <Button auto shadow color='gradient' onPress={register}>
                  Sign Up
               </Button>
            </Modal.Footer>
         </Modal>
      </div>
   );
};