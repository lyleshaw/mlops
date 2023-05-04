import type { NextPage } from 'next';
import { Nav } from '../components/navbar/navbar';
import { Layout } from '../components/navbar/layout';
import { Hero } from '../components/hero';
import { Box } from '../components/styles/box';
import { Faq } from '../components/faq';
import { Footer } from '../components/footer';

const Home: NextPage = () => {
    return (
        <Layout>
            <Nav />
            <Box as="main">
                <Hero />
                <Faq />
                <Footer />
            </Box>
        </Layout>
    );
};

export default Home;
