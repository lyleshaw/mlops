import { Col, Row, User, Text, Tooltip } from '@nextui-org/react';
import { EyeIcon } from '../icons/table/eye-icon';
import { IconButton, StyledBadge } from './table.styled';
import { UpdateModel } from '../updateModel';
import { DeleteModel } from '../deleteModel';

interface Props {
   setReloading: any;
   model: any;
   columnKey: string | React.Key;
}

export const RenderCell = ({ setReloading, model, columnKey }: Props) => {
   const handleView = (id: number, port: number) => {
      window.open(`http://101.42.224.144:${port}`, '_blank');
   };

   const handleDelete = (id: number) => {
      console.log('View model', id);
   };

   // @ts-ignore
   const cellValue = model[columnKey];
   const parseISO = (isoString: string): string => {
      const date = new Date(isoString);
      return date.toLocaleDateString() + " " + date.toLocaleTimeString();
   };
   switch (columnKey) {
      case 'avatar':
         return (
            <User name="" squared src={model.avatar} css={{ p: 0 }}>
            </User>
         );
      case 'status':
         return (
            // @ts-ignore
            <StyledBadge type={String(model.app_status)}>{cellValue}</StyledBadge>
         );
      case 'create':
         return <Text h6>{parseISO(cellValue)}</Text>;
      case 'update':
         return <Text h6>{parseISO(cellValue)}</Text>;
      case 'actions':
         return (
            <Row
               justify="center"
               align="center"
               css={{ 'gap': '$8', '@md': { gap: 0 } }}
            >
               <Col css={{ d: 'flex' }}>
                  <Tooltip content="Details">
                     <IconButton
                        onClick={() => handleView(model.id, model.port)}
                     >
                        <EyeIcon size={20} fill="#979797" />
                     </IconButton>
                  </Tooltip>
               </Col>
               <Col css={{ d: 'flex' }}>
                  <Tooltip content="Edit Model">
                     <UpdateModel setReloading={setReloading} modelID={model.id} />
                  </Tooltip>
               </Col>
               <Col css={{ d: 'flex' }}>

                  <Tooltip
                     content="Delete Model"
                     color="error"
                     onClick={() => handleDelete(model.id)}
                  >
                     <DeleteModel setReloading={setReloading} modelID={model.id} />
                  </Tooltip>
               </Col>
            </Row>
         );
      default:
         return <Text h6>{cellValue}</Text>;
   }
};
