import {
Tab, TabList, TabPanel, TabPanels, Tabs
} from '@chakra-ui/react'
import React from "react";
import axios, {AxiosHeaders} from "axios";
import {GroupListTab} from "./group_lists_tab/GroupListTab";
import {GroupListAdminRequestTab} from "./group_lists_tab/GroupListAdminRequestTab";

export function MainScreen() {

    return (
        <Tabs isFitted variant='enclosed'>
            <TabList mb='1em'>
                <Tab>SendLists</Tab>
                <Tab>Admins Of Lists</Tab>
            </TabList>
            <TabPanels>
                <TabPanel>
                    <GroupListTab></GroupListTab>
                </TabPanel>
                <TabPanel>
                    <GroupListAdminRequestTab></GroupListAdminRequestTab>
                </TabPanel>
            </TabPanels>
        </Tabs>
    )
}
