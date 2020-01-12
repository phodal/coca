passport.use(new FacebookStrategy({}, (req: any, accessToken, refreshToken, profile, done) => {
    if (req.user) {
        User.findOne({ facebook: profile.id }, (err, existingUser) => {
            if (err) { return done(err); }
            if (existingUser) {
                req.flash("errors", { msg: "There is already a Facebook account that belongs to you. Sign in with that account or delete it, then link it with your current account." });
                done(err);
            } else {
                User.findById(req.user.id, (err, user: any) => {
                    if (err) { return done(err); }
                    user.facebook = profile.id;
                    user.tokens.push({ kind: "facebook", accessToken });
                    user.profile.name = user.profile.name || `${profile.name.givenName} ${profile.name.familyName}`;
                    user.profile.gender = user.profile.gender || profile._json.gender;
                    user.profile.picture = user.profile.picture || `https://graph.facebook.com/${profile.id}/picture?type=large`;
                    user.save((err: Error) => {
                        req.flash("info", { msg: "Facebook account has been linked." });
                        done(err, user);
                    });
                });
            }
        });
    }
}));
