import React from 'react';
import { Routes, Route, useLocation } from 'react-router-dom';
import { AnimatePresence, motion } from 'framer-motion';
import ASTTree from './pages/ASTPage.tsx';
import LoginPage from './pages/LoginPage';
import AllSourceCodesPage from './pages/UserSourceCodesPage.tsx';

const pageVariants = {
    initial: {
        opacity: 0,
        x: '-50%',
    },
    in: {
        opacity: 1,
        x: '0%',
    },
    out: {
        opacity: 0,
        x: '50%',
    },
};

const pageTransition = {
    type: 'tween',
    ease: 'easeInOut',
};

const AnimApp: React.FC = () => {
    const location = useLocation();
    return (
        <div className="page-container">
            <AnimatePresence>
                <Routes location={location} key={location.pathname}>
                    <Route
                        path="/"
                        element={
                            <motion.div
                                className="page absolute w-full"
                                initial="initial"
                                animate="in"
                                exit="out"
                                variants={pageVariants}
                                transition={pageTransition}
                            >
                                <ASTTree />
                            </motion.div>
                        }
                    />
                    <Route
                        path="/source_codes"
                        element={
                            <motion.div
                                className="page absolute w-full"
                                initial="initial"
                                animate="in"
                                exit="out"
                                variants={pageVariants}
                                transition={pageTransition}
                            >
                                <AllSourceCodesPage />
                            </motion.div>
                        }
                    />
                    <Route
                        path="/login"
                        element={
                            <motion.div
                                className="page absolute w-full"
                                initial="initial"
                                animate="in"
                                exit="out"
                                variants={pageVariants}
                                transition={pageTransition}
                            >
                                <LoginPage />
                            </motion.div>
                        }
                    />
                </Routes>
            </AnimatePresence>
        </div>
    );
};

export default AnimApp;
