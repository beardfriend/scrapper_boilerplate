import React from "react";
import { Box, Flex, Text, Heading, Image } from "@chakra-ui/react";
import { MetaState, NowMetaState } from "../states/meta";
import { useRecoilValue } from "recoil";
import logo from "../../assets/icons/logo.png";
import { BrowserOpenURL } from "../../../wailsjs/runtime/runtime";

function Footer({ hasSideBar }) {
  const meta = useRecoilValue(MetaState);
  const nowMeta = useRecoilValue(NowMetaState);

  return (
    <Box
      w={hasSideBar ? "calc(100% - 280px)" : "100%"}
      h={hasSideBar ? "15rem" : "23rem"}
      paddingX="5rem"
      bg={hasSideBar ? "navy.origin" : "#0B132B"}
      position="absolute"
      bottom="0"
    >
      <Flex justifyContent="space-between" align="center" h="100%" paddingX="5rem">
        <Flex gap="15rem" align="center">
          {!hasSideBar && (
            <Box w="100%">
              <Image src={logo} w="142px" minW="142px" />
            </Box>
          )}

          <Flex direction="column" gap="1rem" textAlign="left">
            <Box>
              <Text
                color="white"
                cursor="pointer"
                onClick={() => {
                  nowMeta?.allRounderKim?.pricingURL && BrowserOpenURL(nowMeta?.allRounderKim?.pricingURL);
                }}
              >
                Pricing
              </Text>
            </Box>
            <Box>
              <Text
                color="white"
                cursor="pointer"
                onClick={() => nowMeta?.allRounderKim?.payURL && BrowserOpenURL(nowMeta?.allRounderKim?.payURL)}
              >
                Pay
              </Text>
            </Box>
            <Box>
              <Text
                color="white"
                cursor="pointer"
                onClick={() => nowMeta?.allRounderKim?.contactURL && BrowserOpenURL(nowMeta?.allRounderKim?.contactURL)}
              >
                Contact
              </Text>
            </Box>
          </Flex>

          <Flex direction="column" gap="1rem" textAlign="left">
            <Box>
              <Text
                color="white"
                cursor="pointer"
                onClick={() => nowMeta?.allRounderKim?.blogURL && BrowserOpenURL(nowMeta?.allRounderKim?.blogURL)}
              >
                Blog
              </Text>
            </Box>
            <Box>
              <Text
                color="white"
                cursor="pointer"
                onClick={() => nowMeta?.allRounderKim?.storeURL && BrowserOpenURL(nowMeta?.allRounderKim?.storeURL)}
              >
                Store
              </Text>
            </Box>
            <Box>
              <Text
                color="white"
                cursor="pointer"
                onClick={() => nowMeta?.allRounderKim?.youtubeURL && BrowserOpenURL(nowMeta?.allRounderKim?.youtubeURL)}
              >
                Youtube
              </Text>
            </Box>
          </Flex>
        </Flex>
        <Flex direction="column">
          <Heading color="white">{meta?.name}</Heading>
          <Flex align="center" gap="0.5rem" mt="1rem">
            <Heading color="#8F8F8F" fontSize="xl">
              Update:
            </Heading>
            <Heading color="white" fontSize="xl">
              최신: {nowMeta?.allRounderKim?.versionText} 이 기기: {meta?.versionText}
            </Heading>
          </Flex>

          <Flex align="center" gap="0.5rem" mt="1rem">
            <Heading color="#8F8F8F" fontSize="xl">
              Powered By:
            </Heading>
            <Heading color="white" fontSize="xl">
              MINDFLEX
            </Heading>
          </Flex>
          <Text
            onClick={() => nowMeta?.allRounderKim?.patchNoteURL && BrowserOpenURL(nowMeta?.allRounderKim?.patchNoteURL)}
            color="white"
            cursor="pointer"
            mt="1rem"
          >
            버전별 패치노트 보기
          </Text>
        </Flex>
      </Flex>
    </Box>
  );
}

export default Footer;
