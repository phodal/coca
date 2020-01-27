package com.compare;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class TsIdentAppTest {
    @Test
    void shouldHandleField() {
        String str = "passport.use(new FacebookStrategy({}, (req: any, accessToken, refreshToken, profile, done) => {\n" +
                "    if (req.user) {\n" +
                "        User.findOne({ facebook: profile.id }, (err, existingUser) => {\n" +
                "            if (err) { return done(err); }\n" +
                "            if (existingUser) {\n" +
                "                req.flash(\"errors\", { msg: \"There is already a Facebook account that belongs to you. Sign in with that account or delete it, then link it with your current account.\" });\n" +
                "                done(err);\n" +
                "            } else {\n" +
                "                User.findById(req.user.id, (err, user: any) => {\n" +
                "                    if (err) { return done(err); }\n" +
                "                    user.facebook = profile.id;\n" +
                "                    user.tokens.push({ kind: \"facebook\", accessToken });\n" +
                "                    user.profile.name = user.profile.name || `${profile.name.givenName} ${profile.name.familyName}`;\n" +
                "                    user.profile.gender = user.profile.gender || profile._json.gender;\n" +
                "                    user.profile.picture = user.profile.picture || `https://graph.facebook.com/${profile.id}/picture?type=large`;\n" +
                "                    user.save((err: Error) => {\n" +
                "                        req.flash(\"info\", { msg: \"Facebook account has been linked.\" });\n" +
                "                        done(err, user);\n" +
                "                    });\n" +
                "                });\n" +
                "            }\n" +
                "        });\n" +
                "    }\n" +
                "}));\n";
        TsIdentApp.processString(str);
    }
}
