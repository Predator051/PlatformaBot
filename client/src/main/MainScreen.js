import {
Tab, TabList, TabPanel, TabPanels, Tabs
} from '@chakra-ui/react'
import React from "react";
import {ChannelsTab} from "./channels_tab/ChannelsTab";
import {ChannelAdminRequestTab} from "./channels_tab/ChannelAdminRequestTab";
import {ChannelsSubscriptionsTab} from "./channels_tab/ChannelSubscriptionsTab";

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
                    <ChannelsTab></ChannelsTab>
                </TabPanel>
                <TabPanel>
                    <ChannelAdminRequestTab></ChannelAdminRequestTab>
                </TabPanel>
                <TabPanel>
                    <ChannelsSubscriptionsTab></ChannelsSubscriptionsTab>
                </TabPanel>
            </TabPanels>
        </Tabs>
    )
}
