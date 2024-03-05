import {
Tab, TabList, TabPanel, TabPanels, Tabs
} from '@chakra-ui/react'
import React from "react";
import {GroupListTab} from "./group_lists_tab/GroupListTab";
import {GroupListAdminRequestTab} from "./group_lists_tab/GroupListAdminRequestTab";
import {GroupListsSubscriptionsTab} from "./group_lists_tab/GroupListsSubscriptionsTab";

export function MainScreen() {

    return (
        <Tabs isFitted variant='enclosed'>
            <TabList mb='1em'>
                <Tab>Send channels</Tab>
                <Tab>Requests to admins of channels</Tab>
                <Tab>Subscriptions to channels</Tab>
            </TabList>
            <TabPanels>
                <TabPanel>
                    <GroupListTab></GroupListTab>
                </TabPanel>
                <TabPanel>
                    <GroupListAdminRequestTab></GroupListAdminRequestTab>
                </TabPanel>
                <TabPanel>
                    <GroupListsSubscriptionsTab></GroupListsSubscriptionsTab>
                </TabPanel>
            </TabPanels>
        </Tabs>
    )
}
