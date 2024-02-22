import {SendPost} from "../../request/Api";
import React, {useEffect, useState} from "react";
import {MdClose, MdDeck, MdDelete, MdDone} from "react-icons/md"
import {
    Box,
    Button, Center,
    Divider, Flex,
    Input,
    InputGroup,
    InputRightElement,
    ListItem, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay,
    OrderedList, Spacer,
    Stack, useDisclosure
} from "@chakra-ui/react";

export function GroupListAdminRequestTab() {
    let [groupListRequest, setGroupListRequest] = useState([])
    let [groupLists, setGroupLists] = useState([])
    // {
    //     "ID": 11,
    //     "ChatID": 306283632,
    //     "GroupListID": 12,
    //     "Username": "LehaCrump",
    //     "FirstName": "Алексей",
    //     "SecondName": "Кукишев"
    // }
    useEffect(() => {
        SendPost('api/group_lists/admin/requests', {}).then(r => {
            setGroupListRequest(r.data.requests)
            setGroupLists(r.data.groupLists)
            console.log(r.data)
        })
    }, []);

    const clickDeleteGroupList = (id) => {
        // SendPost("api/delete/group_lists", {id}).then(value => {
        //     SendPost('api/group_lists', {}).then(r => {
        //         setGroupListRequest(r.data)
        //         console.log(r.data)
        //     })
        // }, error => {})
    }

    return (
        <Box>
            <Stack spacing={3}>
                <OrderedList>
                    {
                        groupListRequest?.map((v) => <ListItem itemID={v.ID} key={v.ID}>
                            <Flex height={10}>
                                User {' '}
                                [{v.Username} {' '}
                                {v.FirstName} {' '}
                                {v.SecondName} {' '}]
                                requests for administration of {' '}
                                {groupLists?.find(gl => gl.ID == v.GroupListID)?.Name}
                                <Spacer/>
                                <Button colorScheme='green' leftIcon={<MdDone/>} mr={1}>
                                    Accept
                                </Button>
                                <Button colorScheme='red' leftIcon={<MdClose/>}>
                                    Decline
                                </Button>
                            </Flex>
                            <Divider/>
                        </ListItem>)
                    }
                </OrderedList>
            </Stack>
        </Box>

    )
}
